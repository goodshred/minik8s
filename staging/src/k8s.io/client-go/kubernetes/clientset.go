package kubernetes

type Interface interface {
	// client-go包里面的代码大部分使用了代码生成
	//AppsV1() appsv1.AppsV1Interface
	//AppsV1beta1() appsv1beta1.AppsV1beta1Interface
	//StorageV1() storagev1.StorageV1Interface
	//SchedulingV1() schedulingv1.SchedulingV1Interface
}
