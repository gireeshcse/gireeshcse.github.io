# Machine Learning

- Grew out of work in AI
- New capability for computers

Usage Examples :

- Database mining

Large datasets from the growth of automation/web.

- Applications can't program by hand.

Autonomous helicopter, NLP, Computer Vision

- Self-customizing programs

Recommendations

## Definition 

* Arthur Samuel(1959)

Field of study that gives computers the ability to learn without being explicitly programmed.

* Tom Mitchell(1998)

A computer program is said to learn from experience E with respect to some task T and some performance measure P, if its performance on T, as measured by P, improves with experience E.

Example: CLassifying emails as spam or not.

Example: playing checkers.

E = the experience of playing many games of checkers

T = the task of playing checkers.

P = the probability that the program will win the next game.

## Algorithms

- Supervised learning
- Unsupervised learning 
- Reinforcement learning 
- Recommender systems

### Supervised Learning

Common type of machine learning program.

In supervised learning, we are given a data set and already know what our correct output should look like, having the idea that there is a relationship between the input and the output.

Supervised learning problems are categorized into "regression" and "classification" problems. In a regression problem, we are trying to predict results within a continuous output, meaning that we are trying to map input variables to some continuous function. In a classification problem, we are instead trying to predict results in a discrete output. In other words, we are trying to map input variables into discrete categories.

Problems:

- Housing price prediction.
    * Given data about the size of houses on the real estate market, try to predict their price. Price as a function of size is a continuous output, so this is a regression problem.
    * Fiting to straight line or qudaratic function.
    * We could turn this example into a classification problem by instead making our output about whether the house "sells for more or less than the asking price." Here we are classifying the houses based on price into two discrete categories.

Right answers are given.
Regression Problem: Predict continuous valued output(price)

- Breast cancer (malignant, benign)

Probality/ classification problem
Discrete valued output(0 or 1) (0,1,2,3)

Here we can use 1 feature or attribute such as tumor size to classify.

Other attibutes we may use are
    - Clump Thickness
    - Uniformity of cell size
    - uniformity of cell shape
    - age

If we have a problem which have the infinite no of features from which we make predictions, an algorithm called **Support Vector Machine**  allows us to deal with an infinite number of features.


- **Regression** - Given a picture of a person, we have to predict their age on the basis of the given picture
- **Classification** - Given a patient with a tumor, we have to predict whether the tumor is malignant or benign.

### Unsupervised Learning

Data does not have any labels. Can we find a find clusters in given data set.

Example: Google News

Clustering algorithms finds different categories the data falls under.

- Market segmentation
- Social Network analysis
- Astronomical data analaysis

Cocktail party problem.

Seperate audios of different speakers

Examples:

* Given a set of news articles found on the web, group them into sets of articles about the same stories.
* Given a database of customer data, automatically discover market segments and group customers into different market segments.

Unsupervised learning allows us to approach problems with little or no idea what our results should look like. We can derive structure from data where we don't necessarily know the effect of the variables.

We can derive this structure by clustering the data based on relationships among the variables in the data.

With unsupervised learning there is no feedback based on the prediction results.

Example:

Clustering: Take a collection of 1,000,000 different genes, and find a way to automatically group these genes into groups that are somehow similar or related by different variables, such as lifespan, location, roles, and so on.

Non-clustering: The "Cocktail Party Algorithm", allows you to find structure in a chaotic environment. (i.e. identifying individual voices and music from a mesh of sounds at a cocktail party).