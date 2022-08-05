# Algorithms

### Binary Search

* LOG<sub>2</sub> n (How many 2 are required to get n. Example: For 16, we need 4 no. of 2's)

* 2 <sup>10</sup> = 1024

* To check a number in 1024, Binary search takes 10 comparisions at most.

* Sorted order list is required.

* For 4 Billion - With 32 Guess, we can find a number exists or not.

```
def binary_search(list,item):
    low = 0
    high = len(list) - 1
    while low <= high:
        mid = (low + high) / 2
        if list[mid] == item:
            return mid
        if list[mid] < item:
            low = mid + 1
        else:
            high = mid - 1
    return None
```

### Time Complexity

Growth or time it takes as input size grows 

Big O - Worst case runtime (no. of operations)

* O(log n) - Binary Search
* O(n) - Linear Search
* O(n log n) - Quick Sort (Fast Sort Algorithm)
* O(n<sup>2</sup>) - Selection Sort
* O(n!) - Travelling Salesperson (Very Slow)
* O(n log n) - Merge Sort (Both the array must be sorted)

### Travelling Sales Person

SalesPerson to travel n cities  - Problem caliculate minimum travelling distance (O(n!))

* 5 Cities - 5! - 120 Combinations
* 6 Cities - 6! - 720 Combinations
* 7 Cities - 7! - 5040 Combinations

### Arrays Vs Linked List

```
            Arrays      Linked List      
Access      Random      Sequential
Reading     O(1)        O(n)
Insertion   O(n)        O(1)
Deletion    O(n)        O(1)
```
Above Insertion/Deletion of element at the beginning 

### Binary Search with Datastructure


```
                Arrays          BST(Binary Search Tree)     
Searching       O(log n)        O(log n)
Insertion       O(n)            O(log n)
Deletion        O(n)            O(log n)
```

O(n) - Linear Time
O(1) - Constant Time

### Selection Sort

```
void swap(int *num1, *num2)
{
    int temp = *num1;
    *num1 = *num2;
    *num2 = temp;
}

void selectionSort(int arr[],int n)
{
    int i,j;
    for(i = 0;i<n-1;i++)
    {
        for(j=i+1;j<n;j++)
        {
            if(arr[j] < arr[i])
            {
                swap(&arr[j],&arr[i]);
            }
        }
    }
}

```

### Divide and Conquer

* Every Recursive function has 2 cases
    - Base Case
    - Recursive Case
* Quick Sort

* Figuring out the base case solves the problem. This should be the simplest possible case. Divide our problem until it becomes the base case.

#### Divide the plot into equal parts.

```
int small(int a, int b)
{
    if(a == 0)
    {
        return b;
    }
    b = b % a;
    return small(b,a);
}
```

#### Quick Sort

```
int partition(int arr[], int low, int high)
{
    int pivot = arr[high];
    int i = low;
    for(int j = low ; j < high ; j++){
        if(arr[j] < pivot){
            swap(&arr[i],&arr[j]);
            i++;
        }
    }
    swap(&arr[i],&arr[high]);
    return i;
}

void quickSort(int arr[],int low, int high)
{
    if(low < high)
    {
        int pi = partition(arr,low,high);
        quickSort(arr,low,pi - 1);
        quickSort(arr,pi + 1, high);
    }
}
```

In worst case, it is O(n<sup>2</sup>). Depends on pivot

### Hash Tables

* Associative Arrays in javascript
* Dictionaries in Python

* **Collision** - 2 Keys assigned to same slot/value.
* **Load Factor** - No. of items in hash table / Total No. of Slots.
* **Resizing** - LF grows - New Slots to be added
    - Rule of Thumb - Resize if LF > 0.7
    - Result in reinserts.
    - Lower LF, fewer collisions
    - Expensive

* Good hash function distributes values in array evenly.
* Bad HF, groups values togther and produces a lot of collisions.
* Really fast search, insert and delete.
* Used for caching data
* Great for catching duplicates.

### Breadth First Search

* Allows us to find the shortest path b/w two things (Path with fewer segments/Nodes)
* Is there a path from Node A to Node B
* Traversing Algorithm
    - Start traversing from selected node and traverse the graph layerwise thus exploring the neighbour nodes and then move towards next level nodes.
* Graph can contain cycles (visited boolean variable to detect these).
* Use a queue to store the nodes.

```
BFS(G,S) // G - Graph S - start node/ source node
    let Q be queue
    Q.enqueue(S) // Inserting S in queue until all its neighbour vertices are marked
    mark s as visited
    while Q is not empty
        //Removing vertex from queue , whose neighbours will be visited now
        v = Q.dequeue()
        // Processing all neighbours of v
        for all neighbours of w of v in graph G
            if w is not visited
                Q.enqueue(w)
                mark w as visited.

```
Time Complexity - O(Vertices+Edges)

* Applications
    - To determine level of each node in a given tree.

### Dijkstra's Algorithm

* Finds the path with the smallest total weight.
* Each edge in the graph has a number associated with it called weights.
* Only works with Dircted Acyclic Graphs(DAG).
* Can't use this, if we have negative-weight edges - Bellman-Ford is used for this.
* Finds a shortest path tree from a single source node, by building a set of nodes that have a minimum distance from the source.

#### Datastructures

* distances(dist) - Array (Source node S to each node in the graph)
    - Initial Values
        - dist(s) = 0
        - for all other nodes , dist[v] == infinity
* Q - Queue of all nodes in the graph(At end it will be empty)
* S - An empty set, to indicate which nodes are visited (At end S will contain all the nodes)

#### Algorithm

```
function Dijkstra(Graph, Source):
    dist[source] := 0 // Distance from the source to the source.
    for each vertex v in Graph:
        if v not equal to source:
            dist[v] := infinity
        add v to Q
    while Q is not empty:
        v := vertex in Q with min dist[v]
        remove v from Q
        for each neighbour u of v:
            distance := dist[v] + length(v,u)
            if distance < dist[u]: 
                dist[u] := distance
    return dist[]
end function
```

### Greedy Algorithm.

* At each step pick the locally optimal solution and in the end we may end up with the globally optimally solutiion.

* Classroom scheduling problem
    - Pick the class that ends the soonest.
    - Pick the class that starts after the previous class and also which ends the soonest.

* Knapsack problem
    - Change denomination problem
    - Greedy theif problem

* Set covering problem (Approximation Algorithm)
    - Radio show which covers maximum states (Select radio stations which overlaps or covers more than 1 state)

Example:

```
cities_needed = set(["SC","HYD","KKD","RJY","BZA","MAS","SBC","TPTY"])
networks = {}
networks["SUN"] = set(["MAS","SBC","TPTY"])
networks["LOCAL"] = set(["KKD","RJY"])
networks["STAR"] = set(["SBC","HYD","SC"])
networks["ZEE"] = set(["BZA","HYD","KKD"])
networks["HATWAY"] = set(["HYD","SC"])
final_networks = set()
while cities_needed:
    best_network = None
    cities_covered = set()
    for network, cities_for_network in network.items():
        covered = cities_needed & cities_for_network
        if len(covered) > len(cities_covered):
            best_network = network
            cities_covered = covered
    cities_needed  -= cities_covered
    final_network.add(best_network)
// SUN,ZEE, LOCAL,STAR
```
Time complexity - O(n<sup>2</sup>) < O(n!)

#### Traveling Salesperson

```
1 City          -> 1 Route
2 Cities        -> 2 start cities * 1 Route for each start = 2 Total Routes
3 Cities        -> 3 Start Cities * 2 Routes = 6 Total Routes
4 Cities        -> 4 Start Cities * 6 Routes = 24 Total Routes
5 Cities        -> 5 Start Cities * 24 Routes = 120 Total Routes
```

* Traveling salesperson and set covering problem: we have to caliculate every possible solution and pick the smallest/shortest one (NP-Complete)
* If problem involves sequence or set and its hard to solve, it might be NP-Complete.
* If we have an NP-Complete problem, our best bet is to use an approximation algorithm.
* Greedy algorithm is easy to write and fast to run, so they make good approximation algorithms.

### Dynamic Programming.

* Solving sub-problems and which builds up to solving the big problem.
* For Knapsack problem, we will start by solving the problem for smaller knapsacks and then workup to solving the original problem.

* CELL[i][j] = Max of 
    - The previous max (value of CELL[i-1][j])
    - current item weight + value of the remaining space(CELL[i-1][j-current item weight])

Travel Itinery

```
Attraction  |   Time    |   Rating  |   0.5 |   1   |   1.5 |   2
Museum      |   1/2 Day |      7    |   7M  |   7M  |   7M  |   7M
WaterPark   |   1/2 Day |      6    |   7M  |  13MW |  13MW |   13MW
Golconda    |   1 Day   |      9    |   7M  |  13MW |  16GM |   22MW
RFC         |   2 Day   |      9    |   7M  |  13MW |  16GM |   22MW
IMAX        |   1/2 Day |      8    |   8M  |  15MI |  21MWI|   24IGM

```

* Dynamic programming only works when each subproblem is discrete- when it does not depend on other sub problems.
* Best solution may not fill the knapsack completely.

Camping Problem

```
Item  | Weight | Value |  1  |   2   |   3  |   4   |   5   |   6
Water |  3     | 10    |  -  |   -   | 10W  |  10W  |  10W  |   10W 
Book  |  1     |  3    |  3B |   3B  | 10W  |  13WB |  13WB |   13WB 
Food  |  2     |  9    |  3B |   9F  | 12BF |  13WB |  19WF |   22WBF 
Jacket|  2     |  5    |  3B |   9F  | 12BF |  14FJ |  19WF |   22WBF 
Camera|  1     |  6    |  6C |   9BC | 15CF |  18BFC|  20FJC|   25WFC 
// Food, Water, Camera will be carried
```
* Dynamic programming is useful when we are trying to optimize something given a constraint.
* In Knapsack, we have to optimize/maximize the value, constrained by size of the knapsack.

* Tips
    - Every DP solution involves a grid
    - The values in the cells are usually what we are trying to optimize. For knapsack problem, the values were the values of the goods.
    - Each cell is a subproblem, so we need to think about how we can divide our problem into subproblem. That will help us figure out what the axes are.

#### Longest common substring.

```
if word_a[i] == word_b[j]:
    cell[i][j] = cell[i-1][j-1] + 1
else:
    cell[i][j] = 0
```

```
    H   I   S   H
F   0   0   0   0
I   0   1   0   0
S   0   0   2   0
H   1   0   0   3
```

```
    V   I   S   T   A
H   0   0   0   0   0
I   0   1   0   0   0
S   0   0   2   0   0
H   0   0   0   0   0
```

The final solution may not be in the last cell.

For Knapsack problem, last cell always had the final solution. But longest common substring, the solution is largest number in the grid. 

#### Longest Common Subsequence

```
    F   O   S   H
F   1   0   0   0
O   0   2   0   0
R   0   0   0   0
T   0   0   0   0
```

```
    F   O   S   H
F   1   1   1   1
I   1   1   1   1
S   1   1   2   2
H   1   1   2   3
```

```
    F   O   S   H
F   1   0   0   0
I   0   0   0   0
S   0   0   1   0
H   0   0   0   2
```

```
if word_a[i] == word_b[j]:
    CELL[i][j] = CELL[i-1][j-1] + 1
else:
    CELL[i][j] = max(CELL[i-1][j],CELL[i][j-1] )
```

* Longest common subsequence is used to find similarities in DNA strands.
* git diff -> string similarity. Levenishtein distance measures how similar 2 strings are 
* Spell check, figuring whether a user is uploading copyrighted data.

### K-nearest neighbour (KNN)

* Classification
* Recommendation systems
* Feature extraction.
    * For fruit classification (Size and color important)
    * Malignant cancer (Size of tumor, uniformity of cell size)
    * Features affect the size or dimensions of graphs.

Distance b/w 2 points in a graph - $\sqrt{(x2-x1)^2+(y2-y1)^2}$

5 Dimensions -  $\sqrt{(a2-a1)^2+(b2-b1)^2+(c2-c1)^2+(d2-d1)^2+(e2-e1)^2}$

In Classification, we are trying to predict results in a discrete output i.e mapping i/p variables into discrete categories.

* Regression, Predict results within a continuous output, i.e mapping i/p variables to some continuous function.

    - Predicting a response (like a number)
    - Useful for bakery
        - How may puffs to make today?
            - Features
                * Weather Today - Scale (1 to 5) (1 = bad , 5 = great)
                * Weekend or holiday (1 - True , 0 - False)
                * Cricket match
                * Past sales

#### Cosine Similarity

* Different people rate differently - we use cosine similarity , to know they have some taste.
* It is a metric used to measure how similar the documents are irrespective of their size.
* It does not measure the distance b/w two vectors, instead it compares the angles of the two vectors.
* Mathematically, it measures the cosine of the angle between two vectors projected in a multi-dimensional space.
* Advantageous because even if the two similar documents are far apart by the euclidean distance(due to size of the document), chances are they may still be oriented closer together. The smaller the angle higher the cosine similarity.

- Two vectors 
    - Arrays containing the word counts of two documents.
* When plotted on a multi-dimensional space, where each dimension corresponds to a word in the document, it captures the orientation(the angle) of the documents and not the magnitude.

Word Frequency

```
        Doc : Sachin        Doc: Dhoni      Doc: Dhoni Small (Subsection)

Dhoni   10                  400             10  
Cricket 50                  100             5
Sachin  200                  20             1
```

Metrics

```
                Euclidean Distance          Cosine Similarity
Sachin & Dhoni  432.4                       0.15
DSmall & Dhoni  204.0                       0.23
DSmall & Sachi  401.85                      0.77
```
### Picking Good Features

For KNN to work, it is really important to pick the right features to compare against.

- Examples
    - Features that correlate to the movies we are trying to recommend
    - Features that won't have a bias

- OCR (Optical Character Recognition)
    - KNN - Lines, points, curve

- Building a spam filter - Naive Bayes Classifier.