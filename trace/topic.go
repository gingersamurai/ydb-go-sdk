package trace

import (
	"context"
)

// tool gtrace used from ./internal/cmd/gtrace

//go:generate gtrace

type (
	// Topic specified trace of topic reader client activity.
	// gtrace:gen
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	Topic struct {
		// TopicReaderCustomerEvents - upper level, on bridge with customer code
		OnReaderStart func(info TopicReaderStartInfo)

		// TopicReaderStreamLifeCycleEvents
		OnReaderReconnect        func(TopicReaderReconnectStartInfo) func(TopicReaderReconnectDoneInfo)
		OnReaderReconnectRequest func(TopicReaderReconnectRequestInfo)

		// TopicReaderPartitionEvents
		OnReaderPartitionReadStartResponse func(
			TopicReaderPartitionReadStartResponseStartInfo,
		) func(
			TopicReaderPartitionReadStartResponseDoneInfo,
		)
		OnReaderPartitionReadStopResponse func(
			TopicReaderPartitionReadStopResponseStartInfo,
		) func(
			TopicReaderPartitionReadStopResponseDoneInfo,
		)

		// TopicReaderStreamEvents
		OnReaderCommit            func(TopicReaderCommitStartInfo) func(TopicReaderCommitDoneInfo)
		OnReaderSendCommitMessage func(TopicReaderSendCommitMessageStartInfo) func(TopicReaderSendCommitMessageDoneInfo)
		OnReaderCommittedNotify   func(TopicReaderCommittedNotifyInfo)
		OnReaderClose             func(TopicReaderCloseStartInfo) func(TopicReaderCloseDoneInfo)
		OnReaderInit              func(TopicReaderInitStartInfo) func(TopicReaderInitDoneInfo)
		OnReaderError             func(TopicReaderErrorInfo)
		OnReaderUpdateToken       func(
			OnReadUpdateTokenStartInfo,
		) func(
			OnReadUpdateTokenMiddleTokenReceivedInfo,
		) func(
			OnReadStreamUpdateTokenDoneInfo,
		)

		// TopicReaderMessageEvents
		OnReaderSentDataRequest     func(TopicReaderSentDataRequestInfo)
		OnReaderReceiveDataResponse func(TopicReaderReceiveDataResponseStartInfo) func(TopicReaderReceiveDataResponseDoneInfo)
		OnReaderReadMessages        func(TopicReaderReadMessagesStartInfo) func(TopicReaderReadMessagesDoneInfo)
		OnReaderUnknownGrpcMessage  func(OnReadUnknownGrpcMessageInfo)

		// TopicWriterStreamLifeCycleEvents
		OnWriterReconnect  func(TopicWriterReconnectStartInfo) func(TopicWriterReconnectDoneInfo)
		OnWriterInitStream func(TopicWriterInitStreamStartInfo) func(TopicWriterInitStreamDoneInfo)
		OnWriterClose      func(TopicWriterCloseStartInfo) func(TopicWriterCloseDoneInfo)

		// TopicWriterStreamEvents
		OnWriterCompressMessages       func(TopicWriterCompressMessagesStartInfo) func(TopicWriterCompressMessagesDoneInfo)
		OnWriterSendMessages           func(TopicWriterSendMessagesStartInfo) func(TopicWriterSendMessagesDoneInfo)
		OnWriterReadUnknownGrpcMessage func(TopicOnWriterReadUnknownGrpcMessageInfo)
	}

	// TopicReaderPartitionReadStartResponseStartInfo
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderPartitionReadStartResponseStartInfo struct {
		ReaderConnectionID string
		PartitionContext   context.Context
		Topic              string
		PartitionID        int64
		PartitionSessionID int64
	}

	// TopicReaderStartInfo
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderStartInfo struct {
		ReaderID int64
		Consumer string
	}

	// TopicReaderPartitionReadStartResponseDoneInfo
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderPartitionReadStartResponseDoneInfo struct {
		ReadOffset   *int64
		CommitOffset *int64
		Error        error
	}

	// TopicReaderPartitionReadStopResponseStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderPartitionReadStopResponseStartInfo struct {
		ReaderConnectionID string
		PartitionContext   context.Context
		Topic              string
		PartitionID        int64
		PartitionSessionID int64
		CommittedOffset    int64
		Graceful           bool
	}

	// TopicReaderPartitionReadStopResponseDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderPartitionReadStopResponseDoneInfo struct {
		Error error
	}

	// TopicReaderSendCommitMessageStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderSendCommitMessageStartInfo struct {
		// ReaderConnectionID string unimplemented yet - need some internal changes
		CommitsInfo TopicReaderStreamSendCommitMessageStartMessageInfo
	}

	// TopicReaderStreamSendCommitMessageStartMessageInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderStreamSendCommitMessageStartMessageInfo interface {
		PartitionIDs() []int64
		PartitionSessionIDs() []int64
	}

	// TopicReaderSendCommitMessageDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderSendCommitMessageDoneInfo struct {
		Error error
	}

	// TopicReaderCommittedNotifyInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderCommittedNotifyInfo struct {
		ReaderConnectionID string
		Topic              string
		PartitionID        int64
		PartitionSessionID int64
		CommittedOffset    int64
	}

	// TopicReaderErrorInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderErrorInfo struct {
		ReaderConnectionID string
		Error              error
	}

	// TopicReaderSentDataRequestInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderSentDataRequestInfo struct {
		ReaderConnectionID       string
		RequestBytes             int
		LocalBufferSizeAfterSent int
	}

	// TopicReaderReceiveDataResponseStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReceiveDataResponseStartInfo struct {
		ReaderConnectionID          string
		LocalBufferSizeAfterReceive int
		DataResponse                TopicReaderDataResponseInfo
	}

	TopicReaderDataResponseInfo interface {
		GetBytesSize() int
		GetPartitionBatchMessagesCounts() (partitionCount, batchCount, messagesCount int)
	}

	// TopicReaderReceiveDataResponseDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReceiveDataResponseDoneInfo struct {
		Error error
	}

	// TopicReaderReadMessagesStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReadMessagesStartInfo struct {
		RequestContext     context.Context
		MinCount           int
		MaxCount           int
		FreeBufferCapacity int
	}

	// TopicReaderReadMessagesDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReadMessagesDoneInfo struct {
		MessagesCount      int
		Topic              string
		PartitionID        int64
		PartitionSessionID int64
		OffsetStart        int64
		OffsetEnd          int64
		FreeBufferCapacity int
		Error              error
	}

	// OnReadUnknownGrpcMessageInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	OnReadUnknownGrpcMessageInfo struct {
		ReaderConnectionID string
		Error              error
	}

	// TopicReaderReconnectStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReconnectStartInfo struct {
		Reason error
	}

	// TopicReaderReconnectDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReconnectDoneInfo struct {
		Error error
	}

	// TopicReaderReconnectRequestInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderReconnectRequestInfo struct {
		Reason  error
		WasSent bool
	}

	// TopicReaderCommitStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderCommitStartInfo struct {
		RequestContext     context.Context
		Topic              string
		PartitionID        int64
		PartitionSessionID int64
		StartOffset        int64
		EndOffset          int64
	}

	// TopicReaderCommitDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderCommitDoneInfo struct {
		Error error
	}

	// TopicReaderCloseStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderCloseStartInfo struct {
		ReaderConnectionID string
		CloseReason        error
	}

	// TopicReaderCloseDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderCloseDoneInfo struct {
		CloseError error
	}

	// TopicReaderInitStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderInitStartInfo struct {
		PreInitReaderConnectionID string
		InitRequestInfo           TopicReadStreamInitRequestInfo
	}

	// TopicReadStreamInitRequestInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReadStreamInitRequestInfo interface {
		GetConsumer() string
		GetTopics() []string
	}

	// TopicReaderInitDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	TopicReaderInitDoneInfo struct {
		ReaderConnectionID string
		Error              error
	}

	// OnReadUpdateTokenStartInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	OnReadUpdateTokenStartInfo struct {
		ReaderConnectionID string
	}

	// OnReadUpdateTokenMiddleTokenReceivedInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	OnReadUpdateTokenMiddleTokenReceivedInfo struct {
		TokenLen int
		Error    error
	}

	// OnReadStreamUpdateTokenDoneInfo
	//
	// Experimental
	//
	// Notice: This API is EXPERIMENTAL and may be changed or removed in a
	// later release.
	OnReadStreamUpdateTokenDoneInfo struct {
		Error error
	}

	////////////
	//////////// TopicWriter
	////////////

	TopicWriterReconnectStartInfo struct {
		WriterInstanceID string
		Topic            string
		ProducerID       string
		Attempt          int
	}

	TopicWriterReconnectDoneInfo struct {
		Error error
	}

	TopicWriterInitStreamStartInfo struct {
		WriterInstanceID string
		Topic            string
		ProducerID       string
	}

	TopicWriterInitStreamDoneInfo struct {
		SessionID string
		Error     error
	}

	TopicWriterCloseStartInfo struct {
		WriterInstanceID string
		Reason           error
	}

	TopicWriterCloseDoneInfo struct {
		Error error
	}

	TopicWriterCompressMessagesStartInfo struct {
		WriterInstanceID string
		SessionID        string
		Codec            int32
		FirstSeqNo       int64
		MessagesCount    int
		Reason           TopicWriterCompressMessagesReason
	}

	TopicWriterCompressMessagesDoneInfo struct {
		Error error
	}

	TopicWriterSendMessagesStartInfo struct {
		WriterInstanceID string
		SessionID        string
		Codec            int32
		FirstSeqNo       int64
		MessagesCount    int
	}

	TopicWriterSendMessagesDoneInfo struct {
		Error error
	}

	TopicOnWriterReadUnknownGrpcMessageInfo struct {
		WriterInstanceID string
		SessionID        string
		Error            error
	}
)

type TopicWriterCompressMessagesReason string

const (
	TopicWriterCompressMessagesReasonCompressData                = TopicWriterCompressMessagesReason("compress-on-send")           //nolint:lll
	TopicWriterCompressMessagesReasonCompressDataOnWriteReadData = TopicWriterCompressMessagesReason("compress-on-call-write")     //nolint:lll
	TopicWriterCompressMessagesReasonCodecsMeasure               = TopicWriterCompressMessagesReason("compress-on-codecs-measure") //nolint:lll
)

func (r TopicWriterCompressMessagesReason) String() string {
	return string(r)
}
