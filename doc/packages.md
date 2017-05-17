# packages

current:

    main (rpg) -> rpg (rpg/rpg)

future:

    main (rpg) -> scenes (rpg/scenes) -> rpg (rpg/rpg)
		  rpg/rpg to be renamed
      one-way access only
      scenes can access engine code and game logic
      engine cannot access scenes or game logic
      game logic cannot access scenes or engine
