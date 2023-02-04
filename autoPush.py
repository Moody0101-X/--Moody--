
from subprocess import run, PIPE
from json import load, dump
from sys import argv
# mom's number: +212 67208-0551
MSG = argv[1]
pid = None
conFile = "./autoConf.json"
with open(conFile) as fp:
    pid = int(load(fp)["prev_commit_id"]) + 1

CMD_STRING = f"git commit -m \"#{pid} {MSG}\""

run("git add -A", shell=True, stdout=PIPE, encoding="utf-8")
c = run(CMD_STRING, shell=True, stdout=PIPE, encoding="utf-8")
if len(argv) == 3:
    if argv[2] == '-p':
        p = run("git push origin main", shell=True, stdout=PIPE, encoding="utf-8")
        if p.returncode == 0:
            print("Pushed!")
        else:
            print("failed to push!")

if c.returncode == 0:
    print(f"#{pid} {MSG}")
    with open(conFile, "w+") as fp:
        dump({
            "prev_commit_id": pid
        }, fp)

