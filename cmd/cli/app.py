import asyncclick as click
from nodriver.core.connection import asyncio
from simulator.bot import Simulator
@click.command()
@click.version_option("0.1.0", prog_name="hello")
async def hello():
    sim = Simulator()
    await sim.create()
    await sim.run()
    click.echo("Browser Running",color=True)

if __name__ == "__main__":
    asyncio.run(hello())
