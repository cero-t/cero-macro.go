# cero-macro.go
Cero-macro is joystick macro runner for Windows applications.
You can edit and run any kind of joystick macros.

## Why cero-macro?
Some kinds of games block simple keyboard macros such as WSH program or AutoHotKey.
But they don't block joystick control using original device driver like vJoy.

vJoy : http://vjoystick.sourceforge.net/site/

Cero-macro uses vJoy SDK, or Cero-macro is a kind of front-end for vJoy SDK.

## How to use?

1. First of all, setup vJoy v2.0 or lator.
2. Download macro.exe from download link.
3. Just run macro.exe.
4. Open http://localhost:8080/ in any browser.
5. Input some macros in the textboxes of "player1" and "player2".
 * Both or one of the textboxes should be filled.
6. Push "submit" button then wait and see!

## Macro references (currently for fighting game only)

1. One line should have "command" and "frame".
 * The "command" is input commands combined by comma. : cf) 2,lp,mp
 * The "frame" is number of frames. One frame equals to 1/60 sec. : cf) 12
 * The "command" and "frame" are combined by whitespaces. : cf ) 2,lp,mp 12
2. Command reference
 * lp : light punch or jab
 * mp : middle punch or strong
 * hp : hard punch or fierce
 * lk : light kick or short
 * mk : middle kick or forward
 * hk : hard kick or roundhouse
 * 1-9 : direction of 10-keys. 1 is down left, 9 is upper right.

## Macro examples

1. fireball (light)
 * 2 2
 * 3 2
 * 6,hp 2
2. dragon punch (hard)
 * 6 2
 * 2 2
 * 3,hp 2
3. super fireball
 * 2 2
 * 3 2
 * 6 2
 * 2 2
 * 3 2
 * 6,hp 2


## Notification

Due to frameskips, commands can be delayed or skipped sometimes.
Then any single command should have more than or equal to 2 frames for stable running.
