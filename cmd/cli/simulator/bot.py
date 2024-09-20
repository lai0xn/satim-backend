import nodriver as uc
import redis
import asyncio
from .card  import Card
class Simulator:
    def __init__(self) -> None:
        self.redis = redis.Redis(host="localhost",port=6379,db=0)

    async def create(self):
         self.browser = await uc.Browser.create(browser_args=["--proxy-server=http://localhost:8000"])


    async def login(self):
        page = await self.browser.get('https://naviguih.com/SignIn')
        await asyncio.sleep(3)
        username_input = await page.select('input[id="email"]')
        password_input = await page.select('input[id="password"]')
    
        await username_input.send_keys('anis.cheklat@satim.dz')   # Fill the username
        await password_input.send_keys('Anis.Cheklat123@')   # Fill the password
    
        await asyncio.sleep(1)


        submit_btn = await page.select("button")
        await submit_btn.click()
        await asyncio.sleep(2)
    

    async def checkout(self):

        page = await self.browser.get("https://naviguih.com/plans")
        plan_btn = await page.find("Get started")
        await plan_btn.click()
        await asyncio.sleep(2)
        check = await page.select("input[type=checkbox]")
        await check.click()
        bank = await page.select("input[id=baridimob-radio]")
        await asyncio.sleep(3)
        await bank.click()
        terms = await page.select("input[id=link-radio]")
        await asyncio.sleep(3)
        await terms.click()
        checkout = await page.find('Checkout')
        await asyncio.sleep(5)
        await checkout.click()
        await asyncio.sleep(5)
        cib = Card("6280580610061011","341","01 - January","2025","123456")
        await cib.fill_card(page)
        await asyncio.sleep(5)
        await self.payment_confirm(cib,page)
        input()


    async def payment_confirm(self,card,page:uc.Tab):
        password = await page.select("input[type=password]")
        await password.send_keys(card.password)
        btn = await page.find("Valider",best_match=True)
        await btn.click()
        input()


    async def run(self):
        await self.create()
        await self.login()
        await self.checkout()
