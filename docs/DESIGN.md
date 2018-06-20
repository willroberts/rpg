# Design

This game is a roguelike RPG heavily inspired by classic PC games
like [Castle of the Winds](https://en.wikipedia.org/wiki/Castle_of_the_Winds).

Castle of the Winds is tile-based or grid-based roguelike RPG in
which the player character starts in a town and ventures out into a
premade (not randomized) dungeon.

Castle's engine uses a step-based flow for processing events. Each
time the player takes a step, the engine advances one frame. This
advancement affects player location, enemy location, combat events,
vision, and more. However, this approach also keep game flow logic
extremely simple: there are no real-time events to process.

The camera perspective in Castle of the Winds is orthogonal, in
contrast to isometric games of this kind such as Tactics Ogre.

Additionally, Castle of the Winds has no music or sound effects.

This game intends to mimic nearly every design aspect of Castle of
the Winds, with the following exceptions and variations:

## Randomly Generated Content

The game will start in the "starting town", which will be a static
area. The starting town will contain a portal to a random dungeon
floor based on the character's level. Completing a level will return
the character to town, where they have an opportunity to interact
with vendors and tweak their equipment before entering the next
floor.

## Animations, Sounds, and Music

If all other features are completed and the game is generally
playable, the game should include some basic animations, sounds, and
music.

## Score Tracking

Castle of the Winds tracked player score on the "game over" screen,
which was a tombstone with the player's name on it. This game may
take high scores up a notch with real-time leaderboards hosted on a
public web server, though this may be trickier with open source code
for reporting scores. :)
