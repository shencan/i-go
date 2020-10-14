package greedy

// 贪心算法
/*
贪心算法：当下做局部最优，从而达到全局最优（实际情况是每一步都最优可能结果也不是最优）
回溯：能够回退
动态规划：最优判断+回退（保存当前的结果并在适当的时候进行回退）（带最优判断的回溯就叫动态规划）

贪心算法可以解决一些最优化问题，如：求图中的最小生成树、求哈夫曼编码等。然而对于工程和生活中的一些问题，贪心法一般不能得到我们所要求的答案。

一旦某个问题可以通过贪心算法来解决，那么贪心算法一般就是解决该问题的最好办法。
由于贪心法的高效性以及其所得答案比较接近最优结果，贪心法也可以用作辅助算法或者直接解决一些要求结果不特别精确的问题。
*/