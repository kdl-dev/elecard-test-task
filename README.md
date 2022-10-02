# elecard-test-task

---
### Task description
<a>http://contest.elecard.ru/</a>

---
### Instruction for starting

1. <b>```docker build . -t elecard-test-task```</b>
2. <b>```docker run --rm  elecard-test-task```</b> 

---

### CLI methods

1. <b>```docker run --rm  elecard-test-task -m AutoExec```</b> - for automatic execution of all tasks (it is not necessary to specify a method for automatic execution).
2. <b>```docker run --rm  elecard-test-task -m GetTasks```</b> - method for getting the coordinates of circles from the server.
3. <b>```docker run --rm  elecard-test-task -m CheckResults -p 0,0,0,0,..5,4```</b> -  method to check the coordinates of a rectangle. To specify the coordinates, you must specify them separated by a comma after the "p" parameter in the following order: left_bottom_x1, left_bottom_y1, right_top_x1, right_top_y1, left_bottom_x2, etc.. 
---