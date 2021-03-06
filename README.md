[![Build Status](https://travis-ci.org/LinMAD/gobattle.svg?branch=master)](https://travis-ci.org/LinMAD/gobattle)

```text
Incident Report: 3.141592
Date: ██/██/██
Location: Skagerrak; coordinates ██/██/██
Entry: After ██████, our ██████ supercarrier "Govern". 
We had been monitoring working capacity around 2 months and now we loosed our communication with "Govern".
Our emergency fleet ████████  was moved to escalate issue but uncertainty they all gone.

2:42 PM: ████ reaching coordinate ██/██/██, still no visual confirmation.
2:50 PM: ████ I retrieving strange radio transmission on our █ ███ Ghz frequencies. 
The message contains repeating numbers like "0011110000-0000000111" its mismatch our standards.
2:55 PM: ████ I have visual on "Govern", radio transmission changed... they say... HQ we under attack...
2:56 PM: The "Govern" approx. 20 meters away from destroyer "Corax" and it continues attacking.
2:57 PM  We have heavy casualties almost our fleet destroyed... they everywhere...
2:59 PM: HQ, we heavily damaged... "Govern" intercepted control of our drone boats...

That was the last transmission from our emergency fleet ████. We cannot allow another incident such as this.
```

##### Objective: 
`You should implement better AI to eliminate "Govern" fleet.`

##### Rules: 
- You must provide coordinates on `X, Y axis` to fire  

Here is your fleet, each ship by one
- [][][][][][]
- [][][][][]
- [][][][]
- [][][]
- [][]
- []

Example on battlefield:
```text
Player it's your fleet
Y
9|  .  .  .  .  .  .  .  .  .  . 
8|  .  #  .  .  .  .  .  .  #  . 
7|  .  #  .  .  .  .  .  .  .  . 
6|  .  #  .  .  .  .  .  #  .  . 
5|  .  .  .  .  #  .  .  #  .  . 
4|  .  .  .  .  .  .  .  .  .  . 
3|  #  #  #  #  #  .  .  .  .  . 
2|  .  .  .  .  .  .  .  .  .  . 
1|  .  .  .  .  .  .  .  .  .  . 
0|  #  #  #  #  #  #  .  .  .  . 
------------------------------
  X 0  1  2  3  4  5  6  7  8  9 
g
Battlefield of Player               Battlefield of Govern
Y                                    Y
9|  *  .  .  .  .  .  .  .  .  .     9|  .  .  .  .  .  .  .  .  .  . 
8|  X  .  .  .  .  .  .  .  .  .     8|  .  .  .  .  *  .  .  .  .  . 
7|  X  .  .  .  .  *  .  .  .  .     7|  .  .  .  .  *  .  .  .  .  . 
6|  X  .  .  .  .  X  *  .  .  .     6|  *  X  *  .  .  .  .  .  .  . 
5|  X  .  .  .  .  X  .  .  .  .     5|  .  .  .  .  .  .  *  X  *  . 
4|  *  .  .  *  *  *  .  .  .  .     4|  .  .  *  .  .  .  .  .  .  . 
3|  .  *  X  X  X  X  X  *  .  .     3|  X  X  X  X  X  *  .  .  *  . 
2|  *  .  *  .  .  .  .  .  .  .     2|  .  *  .  .  .  .  .  .  .  . 
1|  .  *  .  .  .  .  .  .  .  .     1|  .  .  .  *  .  .  .  *  .  . 
0|  *  .  .  .  .  .  .  .  .  .     0|  X  X  X  X  X  X  *  .  .  . 
------------------------------       ------------------------------
  X 0  1  2  3  4  5  6  7  8  9      X 0  1  2  3  4  5  6  7  8  9 


Enter coordinate to fire:
Target X coordinate: 1
Target Y coordinate: 2

--- GAME END ---
--- Player defeated, whole fleet destroyed--- 

```
