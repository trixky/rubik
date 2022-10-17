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
Even better link https://cube20.org/src/cubepos.pdf<br>
For id rules https://forum.francocube.com/viewtopic.php?t=5952<br>
For cube's group identities :<br>
https://www.jaapsch.net/puzzles/thistle.htm<br>
https://forum.francocube.com/viewtopic.php?t=5952<br>
fig.3 is cool https://arxiv.org/pdf/1601.05744.pdf<br>
http://www.diva-portal.org/smash/get/diva2:816583/FULLTEXT01.pdf<br>
https://puzzling.stackexchange.com/questions/5402/what-is-the-meaning-of-a-tetrad-twist-in-thistlethwaites-algorithm<br>