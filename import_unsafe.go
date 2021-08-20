package kvstore

// UnsafeApplier 作用是给予异步操作能力的
// 其实现目的在于支持先 MockApply 来获得修改后的结果
// 如果满足一定的条件以后执行 ReplaceInner 进行应用，如果不满足则放弃
// 然而如果使用 Applier 中的版本，会造成 MockApply 和 ReplaceInner 之间可能会出现修改，导致 ReplaceInner 执行后中间修改丢失
// 因此使用 UnsafeApplier 的流程为
// - UnsafeMockApplyAndLock 获得修改结果，检查
// - 如果检查通过，利用 UnsafeReplaceInnerWithoutLock 和 UnsafeLock 应用
// - 如果检查失败，利用 UnsafeUnlock 放弃
type UnsafeApplier interface {
	UnsafeLocker
	Applier

	UnsafeApply(actions []*Action) error
	UnsafeApplyAction(action *Action) error

	UnsafeMockApplyAndLock(actions []*Action) (Store, error)
	UnsafeReplaceInnerWithoutLock(new *Map)
}
