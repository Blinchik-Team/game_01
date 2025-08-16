-- Модуль для управления состояниями игры
local M = {}

-- Возможные состояния: "main_menu" (главное меню), "game" (игра активна), "game_pause" (игра на паузе)
local MAIN_MENU = "main_menu"
local GAME = "game"
local GAME_PAUSE = "game_pause"
local GAME_OVER = "game_over"

M.state = MAIN_MENU

-- Функция для установки нового состояния игры
function M.set_state(new_state)
	-- Если устанавливается состояние "game_pause" (пауза)
	if new_state == GAME_PAUSE then
		-- Отправляем сообщение в коллекцию объекту "main:/handler" для установки временного шага в 0 (пауза физики/анимации)
		-- factor = 0 останавливает обновление, mode = 1 указывает на непрерывный режим
		msg.post("main:/handler", "set_time_step", { factor = 0, mode = 1 })
		-- Выводим отладочное сообщение с новым состоянием
		print("GAME_STATE: " .. new_state)
		-- Если устанавливается состояние "game" (игра активна)
	elseif new_state == GAME then
		-- Отправляем сообщение для установки временного шага в 1 (нормальное обновление игры)
		msg.post("main:/handler", "set_time_step", { factor = 1, mode = 1 })
		-- Выводим отладочное сообщение с новым состоянием
		print("GAME_STATE: " .. new_state)
		-- Если устанавливается состояние "main:/handler" (главное меню)
	elseif new_state == MAIN_MENU then
		-- Никаких дополнительных сообщений не отправляем, только логируем
		print("GAME_STATE: " .. new_state)
	elseif new_state == GAME_OVER then
		print("GAME_STATE: " .. new_state)
		msg.post("main:/handler", "set_time_step", { factor = 0, mode = 1 })
	end
	-- Обновляем текущее состояние модуля
	M.state = new_state
end

-- Функция для получения текущего состояния игры
function M.get_state()
	-- Возвращаем текущее значение M.state
	return M.state
end

-- Функция проверки, находится ли игра в состоянии "main_menu"
function M.is_main_menu()
	-- Возвращает true, если текущее состояние — "main_menu", иначе false
	return M.state == MAIN_MENU
end

-- Функция проверки, находится ли игра в состоянии "game"
function M.is_playing()
	-- Возвращает true, если текущее состояние — "game", иначе false
	return M.state == GAME
end

-- Функция проверки, находится ли игра на паузе
function M.is_paused()
	-- Возвращает true, если текущее состояние — "pause", иначе false
	return M.state == GAME_PAUSE
end

-- Функция для перехода в главное меню
function M.to_main_menu()
	-- Устанавливаем состояние "main_menu" через set_state
	M.set_state(MAIN_MENU)
end

-- Функция для начала игры
function M.start_game()
	-- Устанавливаем состояние "game" через set_state
	M.set_state(GAME)
end

-- Функция для постановки игры на паузу
function M.pause()
	-- Проверяем, что игра в состоянии "game", чтобы избежать паузы из других состояний
	if M.is_playing() then
		-- Устанавливаем состояние "pause" через set_state
		M.set_state(GAME_PAUSE)
	end
end

-- Функция для возобновления игры
function M.resume()
	-- Проверяем, что игра в состоянии "pause", чтобы возобновить только из паузы
	if M.is_paused() then
		-- Устанавливаем состояние "game" через set_state
		M.set_state(GAME)
	end
end

-- Функция для переключения между паузой и игрой
function M.toggle_pause()
	-- Если игра на паузе, возобновляем её
	if M.is_paused() then
		M.resume()
		-- Если игра активна, ставим на паузу
	elseif M.is_playing() then
		M.pause()
	end
end

-- Возвращаем модуль для использования в других скриптах
return M