
关于 option 配置项，这里做一下简单解析:

当使用 retry-go 库时，你可以定义一个重试策略，以决定何时进行重试，以及如何进行重试。以下是一些可以在重试策略中配置的选项：

* retry.Delay(delay time.Duration) ：设置重试之间的延迟时间。可以使用 time.Duration 类型的值，例如 10 * time.Millisecond。
* retry.DelayType(delayType retry.DelayType) ：设置重试延迟的类型。可以是 retry.FixedDelay（等待固定的延迟时间）或 retry.BackOffDelay（在每次重试之后加倍延迟时间）。
* retry.MaxJitter(maxJitter time.Duration) ：设置重试延迟的最大抖动时间。可以使用 time.Duration 类型的值，例如 2 * time.Millisecond。默认为 0。
* retry.MaxDelay(maxDelay time.Duration) ：设置重试延迟的最大时间。可以使用 time.Duration 类型的值，例如 5 * time.Second。默认为 0，表示没有限制。
* retry.Attempts(attempts uint) ：设置重试次数的最大值。默认为 0，表示没有限制。
  retry.RetryIf(retryIfFunc retry.RetryIfFunc) ：设置一个函数，用于确定是否应该重试。该函数接收一个 error 类型的参数，并返回一个布尔值，表示是否应该重试。默认为 nil，表示始终重试。
* retry.LastErrorOnly(lastErrorOnly bool) ：设置是否只记录最后一次错误。如果设置为 true，则只记录最后一次错误，否则记录所有错误。默认为 false。
* retry.OnRetry(onRetryFunc retry.OnRetryFunc) ：设置一个函数，在每次重试之前调用。该函数接收一个 uint 类型的参数，表示已经重试的次数。默认为 nil