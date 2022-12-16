#!/usr/bin/python
# Copyright 2022 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os
import re
import sys
from map import package_map


class Transfer(object):
    # The hertz package has been imported. The format is [{name: ${package name}, alias: ${alias name}}]
    import_hertz_library = []
    # The other package has been imported
    import_other_library = ""
    # hertz package no need to convert
    import_unuse_hertz_library = {}
    import_new_hertz_library = {}

    # backup variables name for context.Context
    backup_var = ["c", "ctx", "context_Context"]

    route_handler_name = []

    before_import = ""
    body = ""
    package_map = {}
    other_package_alias = []

    def __init__(self, rootname):
        self.rootname = rootname

    def format_code(self):
        # print("cd " + self.rootname + " && go fmt ./...")
        os.system("cd " + self.rootname + " && go fmt ./...")
        os.system("cd " + self.rootname + " && go get -t github.com/cloudwego/hertz")
        os.system("cd " + self.rootname + " && go mod tidy && go mod verify")

    def items_dir(self):
        """find all go files in rootname"""
        for main_dir, _, file_name_list in os.walk(self.rootname):
            for file in file_name_list:
                if file.endswith(".go"):
                    yield os.path.join(main_dir, file)

    def handle_alias(self, line, alias) -> None:
        """get a alias from a import line"""
        pattern = re.compile('"(.*)"')
        ele = line.split(" ")
        if len(ele) == 1:
            str_re1 = pattern.findall(line)
            if len(str_re1) == 0:
                return
            package_name_list = str_re1[0].split("/")
            alias.append(
                {
                    "name": str_re1[0],
                    "alias": package_name_list[len(package_name_list) - 1],
                }
            )
            self.import_unuse_hertz_library[
                package_name_list[len(package_name_list) - 1]
            ] = str_re1[0]
        else:
            str_re1 = pattern.findall(line)
            alias.append({"name": str_re1[0], "alias": ele[0]})
            self.import_unuse_hertz_library[ele[0]] = str_re1[0]

    def check_import(self, line: str) -> bool:
        """Check if this import line is the library should to be converted"""
        for key in package_map.keys():
            if line.find(key) != -1:
                return True

        return False

    def handle_import(self, fd):
        """convert import"""
        line = fd.readline()
        has_other_framework = False
        pattern = re.compile('"(.*)"')
        while line:
            if line.startswith("//"):
                self.before_import += line
                line = fd.readline()
            elif re.search(r"import \(", line):
                line = fd.readline()
                while not re.search(r"\)", line):
                    line = line.strip()
                    if self.check_import(line):
                        has_other_framework = True
                        self.handle_alias(line, self.import_hertz_library)
                    else:
                        self.import_other_library += "\n" + line
                        ele = line.split(" ")
                        if len(ele) > 1:
                            self.other_package_alias.append({"alias": ele[0]})
                        else:
                            pattern = re.compile('"(.*)"')
                            str_re1 = pattern.findall(line)
                            if len(str_re1) > 0:
                                packageNameList = str_re1[0].split("/")
                                self.other_package_alias.append(
                                    {"alias": packageNameList[len(packageNameList) - 1]}
                                )
                    line = fd.readline()
                break
            elif re.search("import ", line):
                if self.check_import(line):
                    has_other_framework = True
                    ele = line.split(" ")
                    if len(ele) == 2:
                        str_re1 = pattern.findall(line)
                        self.import_hertz_library.append(
                            {"name": str_re1[0], "alias": "fasthttp"}
                        )
                    else:
                        str_re1 = pattern.findall(line)
                        self.import_hertz_library.append(
                            {"name": str_re1[0], "alias": ele[1]}
                        )
                break
            else:
                self.before_import += line
                line = fd.readline()

        return has_other_framework

    def handle_body(self, fd):
        """convert body"""
        for item in self.import_hertz_library:
            if item.get("name"):
                self.package_map[item["alias"]] = package_map.get(item["name"])
        line = fd.readline()

        while line:
            for key in self.package_map:
                allStructs = self.package_map.get(key)
                if not allStructs:
                    continue
                if key in self.import_unuse_hertz_library:
                    self.import_unuse_hertz_library.pop(key)
                for item in allStructs:
                    # match xxx.yyyy || *xxx.yyyy || \txxx.yyyy || (xxx.yyyy || []xxx.yyyy
                    if re.search(
                        r"[^a-zA-Z]{}\.{}[^a-zA-Z]".format(key, item["name"]), line
                    ):
                        line = re.sub(
                            r"{}\.{}".format(key, item["name"]),
                            "{}.{}".format(
                                item["alias"],
                                item["afterName"]
                                if "afterName" in item
                                else item["name"],
                            ),
                            line,
                        )
                        self.import_new_hertz_library[item.get("pkgName")] = item.get(
                            "alias"
                        )
            self.body += line
            line = fd.readline()

    def handle_import_conflict(self):
        for key in package_map:
            for item in package_map[key]:
                tmp = item["pkgName"].split("/")
                item["alias"] = tmp[len(tmp) - 1]
        for key in package_map:
            for p in package_map[key]:
                for o in self.other_package_alias:
                    if p.get("alias") == o.get("alias"):
                        p["alias"] = "h" + p["alias"]

    def print_file(self, path):
        hertzLibrary = ""
        for key in self.import_new_hertz_library:
            tmp = key.split("/")
            alias = ""
            if tmp[len(tmp) - 1] != self.import_new_hertz_library.get(key):
                alias = self.import_new_hertz_library.get(key)
            hertzLibrary += '{} "{}"\n'.format(alias, key)
        for key in self.import_unuse_hertz_library:
            tmp = self.import_unuse_hertz_library[key].split("/")
            alias = ""
            if tmp[len(tmp) - 1] != key:
                alias = key
            hertzLibrary += '{} "{}"\n'.format(
                alias, self.import_unuse_hertz_library[key]
            )

        line = "{}\nimport({}\n{}\n)\n{}".format(
            self.before_import, self.import_other_library, hertzLibrary, self.body
        )
        wopen = open(path, "w")
        wopen.write(line)
        wopen.close()

    def find_route_handler(self):
        filePath = self.items_dir()
        for path in filePath:
            fd = open(path, "r")
            line = fd.readline()
            while line:
                if line.strip().startswith("//"):
                    line = fd.readline()
                    continue
                elif re.search(
                    r"(GET|POST|DELETE|OPTION|PUT|PATCH|HEAD|Any|USE)\(", line
                ):
                    line = line.strip()
                    items = line.split(",")
                    last_item = items[-1]
                    tmp = re.findall(r"[^\W0-9]\w*", last_item)
                    if len(tmp) != 0:
                        self.route_handler_name.append(tmp[-1])
                line = fd.readline()

    def find_func(self, fd, line, lines):
        var_name = []
        tmp = []
        tmp.append(line)
        line = line.strip().split("(")
        app_ctx_name = re.findall(r"[^\W0-9]\w*", line[-1])[0]
        var_name.append(app_ctx_name)
        line = fd.readline()
        brace_num = 0
        while line:
            temp = re.findall(r"\{", line)
            brace_num = brace_num + len(temp)
            temp = re.findall(r"\}", line)
            brace_num = brace_num - len(temp)
            tmp.append(line)
            if re.search(r"[^a-zA-Z]var [^\W0-9]\w*", line):
                var = re.findall(r"[^\W0-9]\w*", line.strip())
                var_name.extend(var[1:-1])
            elif re.search(":=", line):
                line = line.strip().split(":=")
                var = re.findall(r"[^\W0-9]\w*", line[0])
                var_name.extend(var)
            if brace_num == -1:
                break
            line = fd.readline()
        ctx_name = ""
        for name in self.backup_var:
            if name not in var_name:
                ctx_name = name
                break

        first_line = tmp[0]
        first_line = first_line.split("(")
        for idx, val in enumerate(first_line):
            if idx == len(first_line) - 1:
                lines = lines + "(" + ctx_name + " context.Context, " + val
            elif idx != 0:
                lines = lines + "(" + val
            else:
                lines = lines + val
        for val in tmp[1:]:
            if re.search(r"[^a-zA-Z]{}.Context\(\)".format(app_ctx_name), val):
                val = re.sub(r"{}.Context\(\)".format(app_ctx_name), ctx_name, val)
            lines = lines + val
        return lines, line

    def add_import_context(self, path):
        lines = ""
        fd = open(path, "r")
        line = fd.readline()
        while line:
            if re.search(r"import \(", line):
                lines += line
                lines += '"context"\n'
                while not re.search(r"\)", line):
                    line = fd.readline()
                    if re.search(r"\)", line):
                        break
                    if not re.search('"context"', line):
                        lines += line
            elif re.search("import ", line):
                lines += 'import "context"'
                lines += line
            else:
                lines += line
                line = fd.readline()
        wopen = open(path, "w")
        wopen.write(lines)
        wopen.close()

    def handle_context(self):
        """convert other handle function to hertz handle function"""
        self.find_route_handler()
        filePath = self.items_dir()
        for path in filePath:
            lines = ""
            fd = open(path, "r")
            line = fd.readline()
            need_add_import = False
            while line:
                has_found = False
                for item in self.route_handler_name:
                    if re.search(
                        r"func.*{}\(.* \*app.RequestContext\)".format(item), line
                    ):
                        lines, line = self.find_func(fd, line, lines)
                        has_found = True
                        need_add_import = True
                if not has_found:
                    lines = lines + line
                line = fd.readline()
            wopen = open(path, "w")
            wopen.write(lines)
            wopen.close()
            if need_add_import:
                self.add_import_context(path)

    def pre_handle(self):
        # check go.mod
        if not os.path.exists(os.path.join(self.rootname, "go.mod")):
            print(
                "could not found go.mod in %s, please ensure the path is a go module project"
                % self.rootname
            )
            exit(1)

        # check the git repository is dirty
        exec = os.popen("cd " + self.rootname + " && git status -s", "r")
        if exec.read() != "":
            print(
                "The git repository is dirty, please commit or stall all changes and try again"
            )
            exit(1)

    def transfer(self):
        self.pre_handle()
        self.format_code()
        filePath = self.items_dir()
        for path in filePath:
            fd = open(path, "r")
            has_other_framework = self.handle_import(fd)
            if not has_other_framework:
                fd.close()
                self.reset()
                continue
            self.handle_import_conflict()
            self.handle_body(fd)
            self.print_file(path)
            fd.close()
            self.reset()
        self.format_code()
        self.handle_context()
        self.format_code()

    def reset(self):
        self.before_import = ""
        self.import_hertz_library = []
        self.import_other_library = ""
        self.import_unuse_hertz_library = {}
        self.import_new_hertz_library = {}
        self.body = ""
        self.package_map = {}
        self.other_package_alias = []


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("No project path passed in")
        exit(1)

    transfer = Transfer(sys.argv[1])
    transfer.transfer()
