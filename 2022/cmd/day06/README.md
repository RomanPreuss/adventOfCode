# Day06

Find size of directories

Three types of tasks:
-[x] Create tree containing all files and directories
-[ ] Get size of each folder
-[ ] Find all folders smaller than 100000
    - start bottom up and cancel the branch once it's larger than 100000


            a
        /       \
        b       f3(30)
    /       \
  f1(10)    f2(20)

# Ideas
1. recursively iterate through tree (depth-first) and 