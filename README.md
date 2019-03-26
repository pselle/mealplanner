# Meal Planner

This is a meal planner that will suggest some meals following a plant-based diet so you can batch cook for the week for utmost win.

## How it works

This project mixes main components and a style to createa meal, and gives you how much to put in to get in the neighborhood of solid nutrition. The basic formula is:

base + filling (usually a protein) + style

for example:

rice + lentils + Mexican style

That combination might suggest that the rice have tomato sauce and spices added, maybe get some chipotle in there, add some vegetables like green peppers, for example, that sort of thing. Always add a bunch of vegetables on top of the base mix!!

### Rules

Meals follow the following rules:

* A meal is ~400 calories (400, 3x day is 1200, which covers basic nutrients, on top of which you can snack or drink beer or whatever)
* A meal can be less than 400, but no lower than 300
* A meal can be over 400, but no greater than 450
* The rest of a meal's calories/nutrition are filled with vegetables
* 200-300 calories from a grain/starch, with the rest filled in by a protein
* Inputs are in logical measurements, ex 1/2 cup, not 0.7348 measurements. This is why there are boundaries.

Plans consist of multiple meals, to the desired number of servings (ex. 10 servings is 10 lunches/dinners for the week).

* Meals are 2-3 servings
* Totaling the requested number of meals in a plan

### Inputs

Inputs are stored in a CSV with their name, category (ex. base), calorie values, and quantity (ex. 1, cup), and the dry multiplier if relevant (empty if not).

Data sources are generally USDA based, although when there's a discrepancy between that and the general Google response, I pick the Google response.

### Sample output

To plan 10 meals:

```
> mealplan -m 10

10 meals you say? How about you make:

1. Lentils and brown rice, Italian style, 3 servings
2. Pasta and chickpeas, Mexican style, 3 servings
3. Edamame and quinoa, Asian style, 2 servings
4. Seitan and potatoes, American style, 2 servings

1. Lentils and brown rice, 3 servings
3 cups brown rice, or 1 cup dry
2.5 cups lentils, or 1.25 cup dry
407.66 calories per serving

2. Pasta and chickpeas, 3 servings
...
```

And so on.
