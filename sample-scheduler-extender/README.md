# scheduler-extender

## 介绍extender工作逻辑

### routers.go
我们实现 /filter 和 /prioritize 两个功能的handler function处理程序。Filter 这个扩展函数接收一个输入类型为 schedulerapi.ExtenderArgs 的参数，然后返回一个类型为*schedulerapi.ExtenderFilterResult 的值。

### predicates.go
在过滤函数中，我们循环每个节点然后用我们自己实现的业务逻辑来判断是否应该批准该节点，在podFitsOnNode()函数中我只是简单的**检查随机数是否>=50**来判断即可，如果是的话我们就认为这是一个“被批准”的节点，否则拒绝批准该节点
```go
func podFitsOnNode(pod *v1.Pod, node v1.Node) (bool, []string, error) {

	var failReasons []string

	judge := (rand.Intn(100) >= 50)

	if judge {
		log.Printf("pod %v/%v fits on node %v\n", pod.Name, pod.Namespace, node.Name)

		return true, nil, nil
	}

	log.Printf("pod %v/%v does not fit on node %v\n", pod.Name, pod.Namespace, node.Name)

	failures := "It's not fits on this node."

	failReasons = append(failReasons, failures)

	return false, failReasons, nil

}
```

### priorities.go
在prioritize函数中，我们在每个节点上随机给出一个分数。
```go
func prioritize ( args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
  pod:=args.Pod
  nodes:=args.Nodes.Items
  hostPriorityList := make(schedulerapi.HostPriorityList, len(nodes))
  for i, node := range nodes {
      score := rand.Intn(schedulerapi.MaxPriority+1)// 在最大优先级内随机取一个值               
      log.Printf(luckyPrioMsg,pod.Name,pod.Namespace,score)
      hostPriorityList[i] = schedulerapi.HostPriority{ 
          Host: node.Name,
	  Score: score,
	  }
      }
      return &hostPriorityList
 }
```


