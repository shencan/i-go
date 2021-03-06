package hotrank

/*
公式如下:
Score=(P-1)/(T+2)^G
Score 最终得分
P 即获赞数、投票数、点击数等等
  如果你不想让"高票帖子"与"低票帖子"的差距过大，可以在得票数上加一个小于1的指数，比如(P-1)^0.8。
T 发布到现在过去的时间（小时）用于使得热度随着时间进行衰减
  之所以选择2，可能是因为从原始文章出现在其他网站，到转贴至Hacker News，平均需要两个小时
G 重力因子 控制热度随时间衰减的速度（推荐值1.8）
  G值越大，曲线越陡峭，排名下降得越快，意味着排行榜的更新速度越快。
*/
