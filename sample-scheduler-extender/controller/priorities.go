package controller

import (
	"log"

	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const (
	PrioMsg = "pod %v/%v gets score %v\n"
)


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
