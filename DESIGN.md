# Initial design
Below is a compact outline for a simple incremental (idle) crafting/clicker game with a 2×2 crafting grid that can expand later. The goal is to keep the scope small yet offer a satisfying sense of progression from foraging to mining and smelting.

---

## 1. Overall Progression Outline

1. **Start**:  
   - Manually gather basic resources by picking bushes (for Berries and Fibers) and collecting loose Rocks on the ground.  
   - No tools required for these first actions.  
   - Crafting grid is 2×2 at the beginning.

2. **Early Game** (with a 2×2 Grid):  
   - Craft the first **Stone Axe** or **Stone Pickaxe** out of very simple recipes.  
   - Use those tools to chop small trees for Wood or mine slightly larger rocks for Stone.  
   - Unlock a few additional upgrades and items.

3. **Mid-Game** (expanding the grid):  
   - Acquire more advanced resources: Copper Ore, Iron Ore.  
   - Smelt them into Ingots using a **Furnace**.  
   - Craft better tools (Copper, then Iron) to increase gathering speed.  
   - Earn or craft an item that **expands the crafting grid** to 3×3 (or 3×4) to allow bigger, more complex recipes.

4. **Late(ish) Game** (still small scope):  
   - Automate some resource collection (idle gathering).  
   - Possibly craft advanced structures (like an **Ore Drill** or **Lumberjack Hut**) that passively generate resources.  

---

## 2. Resource List & Gathering

Below is a proposed list of basic resources and how players initially obtain them:

1. **Berries**  
   - Found by picking **Bushes**.  
   - Food item (if you want a small “energy” system or bonus).  
   - Used sparingly in some recipes (e.g., to craft a simple “feed” for an automated gatherer or for a future cooking system).

2. **Fibers**  
   - Also from **Bushes**.  
   - Very early crafting component (replaces “sticks” or “strings”).  
   - Used in simple tools (lash together Rocks to make Stone Tools).

3. **Rocks**  
   - Collected from loose stones on the ground.  
   - Early crafting resource to make stone tools.

4. **Wood**  
   - Gathered by chopping **Small Trees**.  
   - A fundamental building and crafting component.

5. **Stone**  
   - Gathered by mining **Bigger Rocks** or from small quarries once you have a Stone Pickaxe.  
   - Used to craft better tools and structures than those made just with Rocks.

6. **Copper Ore**  
   - Mined from Copper Veins.  
   - Must be **smelted** into Copper Ingots in a Furnace.

7. **Iron Ore**  
   - Mined from Iron Veins (which might require at least a Stone or Copper Pickaxe).  
   - Also smelted into Iron Ingots in a Furnace.

8. **(Refined) Planks**  
   - Process Wood with a **Saw** or a simple recipe in the 2×2 grid (e.g., 2 Wood → 2 Planks).  
   - Used for building upgrades or better wooden structures.

9. **Copper Ingots**  
   - From smelting Copper Ore.  
   - Used to craft Copper Tools (better speed, or access to higher-tier resources).

10. **Iron Ingots**  
    - From smelting Iron Ore.  
    - Used to craft Iron Tools (even better speed, can access anything below steel or diamond-tier if you want to expand later).

---

## 3. Simple Crafting Recipes (2×2 Grid)

Below are some example recipes that fit into a 2×2 grid. (Note: you can decide the exact arrangement within the squares. It could be top row for main materials, bottom row for Fibers, etc.)

1. **Stone Pickaxe**  
   - **Ingredients**: 2 Rocks + 2 Fibers  
   - **Effect**: Allows mining of bigger rocks for **Stone**.  
   - **Stat Benefit**: Increases gathered Stone per click or per second (if automated).

2. **Stone Axe**  
   - **Ingredients**: 2 Wood + 2 Fibers  
   - **Effect**: Allows chopping small to medium trees more efficiently, generating more Wood each time.

3. **Saw (or Woodcutter’s Table)**  
   - **Ingredients**: 2 Stone + 2 Wood  
   - **Effect**: Used to convert Wood → Planks (or placed in a special slot if you want to keep the grid free).  
   - Could also be a direct 2×2 grid craft that yields Planks as soon as you combine Wood in it.

4. **Furnace**  
   - **Ingredients**: 4 Stone  
   - (You could require the 2×2 to be all Stone for simplicity.)  
   - **Effect**: Allows smelting of Ore into Ingots.  
   - Put Ore in, consume some Wood as fuel or Coal if you introduce it, produce Ingots.  
   - This might be done in a special “Furnace UI” instead of the normal 2×2 grid, but it’s up to you.

5. **Copper Pickaxe**  
   - **Ingredients**: 2 Copper Ingots + 2 Fibers  
   - **Effect**: More efficient mining, can break Iron Ore.  

6. **Copper Axe**  
   - **Ingredients**: 2 Copper Ingots + 2 Planks (handle)  
   - **Effect**: Chops larger trees more quickly, yields more Wood.

7. **Iron Pickaxe**  
   - **Ingredients**: 2 Iron Ingots + 2 Planks  
   - **Effect**: Best gathering speed so far for stone, copper, iron.

8. **Iron Axe**  
   - **Ingredients**: 2 Iron Ingots + 2 Planks  
   - **Effect**: Highest-level wood-chopping speed so far.  

*(You can replicate this pattern for upgrades: e.g., Iron Sword if you want to add monsters, etc.)*

---

## 4. Upgrades & Grid Expansion

1. **Crafting Table Upgrade**  
   - **Ingredients**: A mid-tier recipe—maybe 4 Planks + 1 Copper Ingot.  
   - **Effect**: Expands the 2×2 grid to **3×3**.  
   - This unlocks more complex recipes (like additional structures, automation items, etc.).

2. **Automation Structures** (optional small examples)  
   - **Ore Drill** (3×3 recipe): 4 Iron Ingots + 4 Stone + 1 Copper Pickaxe.  
     - Passively gathers Stone, Copper Ore, or Iron Ore over time.  
   - **Lumberjack Hut** (3×3 recipe): 4 Planks + 2 Stone + 1 Copper Axe + 2 Fibers.  
     - Passively gathers Wood over time.  
   - **Bush Farm** (2×2 or 3×3): 4 Fibers + 4 Planks.  
     - Gradually generates Berries and Fibers.

---

## 5. Core Idle Mechanics

- **Manual Clicking**: At first, you (the player) click “Gather Bushes” or “Gather Rocks.” The yield is small, but it’s enough to craft your first Stone tool.  
- **Tool Efficiency**: Holding better tools increases how much you gather or how fast you gather. In an idle sense, that might mean your “per click” yield is higher or your “gather per second” is higher if you have any passive generation.  
- **Automation**: Once you have the right materials, you build small structures that collect resources for you over time.  
- **Crafting Grid**:  
  - 2×2 at the start.  
  - A key milestone is upgrading to 3×3.  
  - Possibly a further expansion to 3×4 or 4×4 if you want to extend the game.

---

## 6. Example Flow (Step by Step)

1. **Initial**  
   - Click “Pick Bushes” → +1 Berry, +1 Fiber every time.  
   - Click “Gather Rocks” → +1 Rock every time.

2. **First Crafts (2×2)**  
   - Craft Stone Pickaxe (2 Rocks + 2 Fibers).  
   - Now you can click “Mine Stones” (bigger rocks) for +1 Stone, +1 Rock or some variable yield.  
   - Craft Stone Axe (2 Wood + 2 Fibers) once you gather enough Wood from small trees.

3. **Unlock Furnace**  
   - Craft Furnace (4 Stone in the 2×2 grid).  
   - Smelt Copper Ore → Copper Ingots.  
   - Now you can craft Copper Tools.

4. **Build Saw / Woodcutter’s Table**  
   - Convert Wood → Planks.  
   - Planks let you build more advanced items.

5. **Crafting Table Upgrade**  
   - 3×3 grid unlocked.  
   - Build advanced automation structures (Drill, Lumberjack Hut, etc.).  
   - Craft Iron tools for maximum early-game efficiency.

---

## 7. Keeping It “Small-Scale” but Flexible

- You only need a handful of items:
  - **Fibers**, **Berries**, **Rocks**, **Stone**, **Wood**, **Copper Ore**, **Iron Ore**, **Copper Ingots**, **Iron Ingots**, **Planks**.
- You only need a handful of tools:
  - **Stone**, **Copper**, **Iron** (Pickaxe/Axe).
- You only need one major **structure** to process resources (the Furnace) and a saw or table to process Wood → Planks.  
- A couple of “automation structures” can be introduced to give the game that idle/automation flair without bloating the design.

---

## 8. Possible Additions or Variations

- **Coal**: If you want an extra step in smelting (fuel management), you can add Coal as a resource from mining.  
- **Different Tree Types** (e.g., Normal Tree, Pine Tree, Oak Tree) for variety, each offering slightly different wood yields or the same resource but at different rates.  
- **Time-Based Growth**: Bushes and trees might regrow after some idle time.  
- **Limited vs. Unlimited Inventory**: You proposed unlimited inventory, so you can avoid the complexity of storage management.  
- **Quests or Achievements**: “Gather 100 Wood” → earn a special item or bonus to keep the sense of progression.  
- **Expansion**: After the 3×3 grid is unlocked, you could keep going with more advanced metals (Silver, Gold, etc.) if you decide to scale up.

---

### Wrap-Up

This design gives you a straightforward incremental game loop:

1. **Manual gathering** of fiber, berries, and rocks.  
2. **First tools** crafted in a simple 2×2 recipe grid.  
3. **Mining stone and ores**, chopping wood for bigger yields.  
4. **Smelting** to obtain ingots.  
5. **Tool Upgrades** leading to faster gathering.  
6. **Crafting Table Upgrade** to expand the grid, unlocking advanced items.  
7. **Automation** structures to gather resources passively.  

Keep the scope tight with just a few resources and items, so the player experiences a clear beginning, middle, and “transition to bigger things” in your idle/clicker game.
