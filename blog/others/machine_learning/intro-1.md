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

Example: Classifying emails as spam or not.

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

### Model representation

### Linear regression with one variable

* Regression Problem -      Predict real-valued output
* SUpervised Learning-      Given right answers for each example of data.
* Classification Problem-   Discrete-valued outputs

Training sets Notations

* m = No. of training examples
* x's = input variable/features
* y's = output variable/ target variable
* (x,y) - One taining example
* (x<sup>i</sup>, y<sup>i</sup>) - i<sup>th</sup>(row) training example.
* h(hypothesis) maps froms x's to y's (Learning algorithms)
* h<sub>&theta;</sub>(x) = &theta;<sub>0</sub> + &theta;<sub>1</sub>x 

<img src="images/basic_ml.png" alt="Machine Learing basic">

#### Cost function

h<sub>&theta;</sub>(x) = &theta;<sub>0</sub> + &theta;<sub>1</sub>x 

&theta;<sub>i</sub>'s : Parameters

Choose &theta;<sub>0</sub>, &theta;<sub>1</sub> so that h<sub>&theta;</sub>(x) is close to y for our training examples (x,y)

Goal : Minimize &theta;<sub>0</sub>, &theta;<sub>1</sub> that is 

J( &theta;<sub>0</sub>, &theta;<sub>1</sub>) = 1/2m * &Sigma; <sup>m</sup><sub>i=1</sub> (h<sub>&theta;</sub>(x) - y)<sup>2</sup>


Minimize J( &theta;<sub>0</sub>, &theta;<sub>1</sub>) is called **cost function** or sometimes called **squared error function** or **Mean squared error** which is mostl commonly used for linear regression problems.

<img src="images/cost_function_linear_regression.png" alt="cost_function_linear_regression">

<img src="images/cost_function_linear_regression1.png" alt="cost_function_linear_regression">

<img src="images/cost_function_linear_regression2.png" alt="cost_function_linear_regression">

<img src="images/cost_function_linear_regression3.png" alt="cost_function_linear_regression">


Final goal, we should try to minimize the cost function. In this case,  &theta;<sub>1</sub>  =1 is our global minimum

**Contour Plots**

Contour plots (sometimes called Level Plots) are a way to show a three-dimensional surface on a two-dimensional plane. It graphs two predictor variables X Y on the y-axis and a response variable Z as contours. These contours are sometimes called z-slices or iso-response values.

A contour plot is appropriate if you want to see how some value Z changes as a function of two inputs, X and Y:
z = f(x,y).

A contour plot is a graph that contains many contour lines. A contour line of a two variable function has a constant value at all points of the same line. An example of such a graph is the one to the right below.

<img src="images/contour1.png" alt="cost_function">

Taking any color and going along the 'circle', one would expect to get the same value of the cost function. For example, the three green points found on the green line above have the same value for J( &theta;<sub>0</sub>, &theta;<sub>1</sub>) and as a result, they are found along the same line. The circled x displays the value of the cost function for the graph on the left when  &theta;<sub>0</sub> = 800 and &theta;<sub>1</sub> = -0.15. Taking another h(x) and plotting its contour plot, one gets the following graphs:

<img src="images/contour2.png" alt="cost_function">


When &theta;<sub>0</sub> = 360 and &theta;<sub>1</sub> =  0, the value of J( &theta;<sub>0</sub>, &theta;<sub>1</sub>) in the contour plot gets closer to the center thus reducing the cost function error. Now giving our hypothesis function a slightly positive slope results in a better fit of the data.

<img src="images/contour3.png" alt="cost_function">

The graph above minimizes the cost function as much as possible and consequently, the result of &theta;<sub>1</sub> and &theta;<sub>0</sub> tend to be around 0.12 and 250 respectively. Plotting those values on our graph to the right seems to put our point in the center of the inner most 'circle'.