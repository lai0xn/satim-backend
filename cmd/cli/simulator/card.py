import random
import asyncio
from nodriver import Tab
from selenium.webdriver.common.keys import Keys

class Card:
    def __init__(self, number, csv, expiration_month, expiration_year,password) -> None:
        self.number = number
        self.csv = csv
        self.expiration_month = expiration_month
        self.expiration_year = expiration_year
        self.password = password

    async def human_delay(self, min_delay=0.5, max_delay=1.5):
        """Introduce a random delay between actions to simulate human behavior."""
        await asyncio.sleep(random.uniform(min_delay, max_delay))

    async def type_like_human(self, input_field, text):
        """Simulate typing one character at a time with random delays."""
        for char in text:
            await input_field.send_keys(char)
            await self.human_delay(0.05, 0.2)  # Small delay between each character
    
    async def move_mouse_and_click(self, element):
        """Simulate human-like mouse movement before clicking."""
        await element.mouse_move()  # Simulate mouse hover before click
        await self.human_delay(0.5, 1.2)  # Random delay before clicking
        await element.mouse_click()

    async def fill_card(self, page: Tab):
        card_input = await page.select("input[id=iPAN]")
        csv_input = await page.select("input[id=iCVC]")
        month = await page.select("select[id=month]")
        year = await page.select("select[id=year]")
        name = await page.select("input[id=iTEXT]")
        submit_btn = await page.select("button[id=buttonPayment]")

        # Debugging: Check if the elements were found
        if not all([card_input, csv_input, month, year, name, submit_btn]):
            raise ValueError("One or more elements could not be found on the page")

        # Simulate human-like typing for the card number
        await self.type_like_human(card_input, self.number)

        # Simulate typing CSV one character at a time
        for digit in self.csv:
            await self.type_like_human(csv_input, digit)
            await self.human_delay()  # Random delay between each digit



        # Simulate typing name
        await self.type_like_human(name, "SATIM Test")
        
        # Random pause before clicking submit
        await self.human_delay(1, 2)

                # Simulate selection of month and year
        await self.move_mouse_and_click(month)
        await self.human_delay()  # Pause after selecting mont    
        await self.type_like_human(month,self.expiration_month)
        await asyncio.sleep(0.5)
        await month.send_keys(Keys.RETURN)
        await self.move_mouse_and_click(year)
        await self.type_like_human(year,self.expiration_year)
        await asyncio.sleep(0.75)
        await month.send_keys(Keys.RETURN)

        # Simulate clicking the submit button
        await self.move_mouse_and_click(submit_btn)


        await asyncio.sleep(15)


