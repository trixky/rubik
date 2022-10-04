# api
## Structure
Seperate corners and edges<br>
Both have an array for position and another for orientation.<br>
For edges, orientation is 0 or 1. For corners, up to 2.<br>

<i>Example</i><br>
Edges:<br>
<pre>
0  1  2  3  4  5  6  7  8  9  10 11
UB UR UF UL FR FL BL BR DF DL DB DR
</pre>

Corners:<br>
<pre>
0   1   2   3   4   5   6   7
ULB URB URF ULF DLF DLB DRB DRF
</pre>

See https://stackoverflow.com/questions/500221/how-would-you-represent-a-rubiks-cube-in-code<br>
Clearer but more math heavy explanation
https://people.math.harvard.edu/~jjchen/docs/Group%20Theory%20and%20the%20Rubik%27s%20Cube.pdf<br>
Even better link https://cube20.org/src/cubepos.pdf
