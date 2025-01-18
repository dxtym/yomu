import os
import sys
import asyncio
import logging

from dotenv import load_dotenv
from aiogram import Bot, Dispatcher
from aiogram.types import Message, InlineKeyboardMarkup, InlineKeyboardButton
from aiogram.filters import CommandStart

load_dotenv()
dp = Dispatcher()

@dp.message(CommandStart())
async def handle_start(message: Message) -> None:
	await message.answer(
		text=f"Ohayo, {message.from_user.full_name}! 👋\nEnjoy the alpha release of yomu. If you find any issues, dm @dxtym 🙏", 
		reply_markup=InlineKeyboardMarkup(inline_keyboard=[
			[InlineKeyboardButton(text="Let's read!", url="t.me/yomubot/yomu")]
		])
	)

async def main() -> None:
	bot = Bot(os.getenv("BOT_TOKEN"))
	await dp.start_polling(bot)

if __name__ == "__main__":
	logging.basicConfig(
		level=logging.INFO,
		stream=sys.stdout
	)
	asyncio.run(main())