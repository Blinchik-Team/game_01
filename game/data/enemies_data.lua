-- /data/enemies_data.lua
local M = {}

M.types = {
	small_random = {
		sprite = "enemy_50",
		size = 50,
		behavior = "random_walker",
		-- Характеристики для Фазы 2
	},
	ranged = {
		sprite = "enemy_100",
		size = 100,
		behavior = "ranged_strafer",
	},
	-- chaser(преследователь)
	chaser_medium = { -- Для врагов 75x75
		sprite = "enemy_75", -- Будет заменяться при спавне
		size = 75,
		behavior = "chaser",
	},
	chaser_big = { -- Для врагов 200x200
		sprite = "enemy_200", -- Будет заменяться при спавне
		size = 200,
		behavior = "chaser",
	},
	chaser_boss = { -- Для врагов 300x300
		sprite = "enemy_300", -- Будет заменяться при спавне
		size = 300,
		behavior = "chaser",
	},
}

return M