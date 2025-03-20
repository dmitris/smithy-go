package http

import "context"

// smithyhttp.Interceptors
type Interceptors struct {
	ReadBeforeExecution             []ReadBeforeExecutionInterceptor
	ModifyBeforeSerialization       []ModifyBeforeSerializationInterceptor
	ReadBeforeSerialization         []ReadBeforeSerializationInterceptor
	ReadAfterSerialization          []ReadAfterSerializationInterceptor
	ModifyBeforeRetryLoop           []ModifyBeforeRetryLoopInterceptor
	ReadBeforeAttempt               []ReadBeforeAttemptInterceptor
	ModifyBeforeSigning             []ModifyBeforeSigningInterceptor
	ReadBeforeSigning               []ReadBeforeSigningInterceptor
	ReadAfterSigning                []ReadAfterSigningInterceptor
	ModifyBeforeTransmit            []ModifyBeforeTransmitInterceptor
	ReadBeforeTransmit              []ReadBeforeTransmitInterceptor
	ReadAfterTransmit               []ReadAfterTransmitInterceptor
	ModifyBeforeDeserialization     []ModifyBeforeDeserializationInterceptor
	ReadBeforeDeserialization       []ReadBeforeDeserializationInterceptor
	ReadAfterDeserialization        []ReadAfterDeserializationInterceptor
	ModifyBeforeAttemptCompletion   []ModifyBeforeAttemptCompletionInterceptor
	ReadAfterAttempt                []ReadAfterAttemptInterceptor
	ModifyBeforeExecutionCompletion []ModifyBeforeExecutionCompletionInterceptor
	ReadAfterExecution              []ReadAfterExecutionInterceptor
}

// smithyhttp.Interceptors
type Interceptors struct {
	BeforeExecution           []BeforeExecutionInterceptor
	BeforeSerialization       []BeforeSerializationInterceptor
	AfterSerialization        []AfterSerializationInterceptor
	BeforeRetryLoop           []BeforeRetryLoopInterceptor
	BeforeAttempt             []BeforeAttemptInterceptor
	BeforeSigning             []BeforeSigningInterceptor
	AfterSigning              []AfterSigningInterceptor
	BeforeTransmit            []BeforeTransmitInterceptor
	AfterTransmit             []AfterTransmitInterceptor
	BeforeDeserialization     []BeforeDeserializationInterceptor
	AfterDeserialization      []AfterDeserializationInterceptor
	BeforeAttemptCompletion   []BeforeAttemptCompletionInterceptor
	AfterAttempt              []AfterAttemptInterceptor
	BeforeExecutionCompletion []BeforeExecutionCompletionInterceptor
	AfterExecution            []AfterExecutionInterceptor
}

// smithyhttp.ReadBeforeExecutionInterceptor
type BeforeExecutionInterceptor interface {
	BeforeExecution(context.Context, *Request, any) error
}
