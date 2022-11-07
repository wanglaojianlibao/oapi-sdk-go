// Package dispatcher code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/vc/v1"
)

// 加入会议
//
// - 发生在有人加入会议时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/join_meeting
func (dispatcher *EventDispatcher) OnP2MeetingJoinMeetingV1(handler func(ctx context.Context, event *larkvc.P2MeetingJoinMeetingV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.join_meeting_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.join_meeting_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.join_meeting_v1"] = larkvc.NewP2MeetingJoinMeetingV1Handler(handler)
	return dispatcher
}

// 离开会议
//
// - 发生在有人离开会议时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/leave_meeting
func (dispatcher *EventDispatcher) OnP2MeetingLeaveMeetingV1(handler func(ctx context.Context, event *larkvc.P2MeetingLeaveMeetingV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.leave_meeting_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.leave_meeting_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.leave_meeting_v1"] = larkvc.NewP2MeetingLeaveMeetingV1Handler(handler)
	return dispatcher
}

// 会议结束
//
// - 发生在会议结束时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/meeting_ended
func (dispatcher *EventDispatcher) OnP2MeetingEndedV1(handler func(ctx context.Context, event *larkvc.P2MeetingEndedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.meeting_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.meeting_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.meeting_ended_v1"] = larkvc.NewP2MeetingEndedV1Handler(handler)
	return dispatcher
}

// 会议开始
//
// - 发生在会议开始时，目前仅提供预约会议的相关事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/meeting_started
func (dispatcher *EventDispatcher) OnP2MeetingStartedV1(handler func(ctx context.Context, event *larkvc.P2MeetingStartedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.meeting_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.meeting_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.meeting_started_v1"] = larkvc.NewP2MeetingStartedV1Handler(handler)
	return dispatcher
}

// 录制停止
//
// - 发生在录制结束时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/recording_ended
func (dispatcher *EventDispatcher) OnP2MeetingRecordingEndedV1(handler func(ctx context.Context, event *larkvc.P2MeetingRecordingEndedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_ended_v1"] = larkvc.NewP2MeetingRecordingEndedV1Handler(handler)
	return dispatcher
}

// 录制完成
//
// - 发生在录制文件上传完毕时
//
// - 收到该事件后，方可进行录制文件获取、授权等操作
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/recording_ready
func (dispatcher *EventDispatcher) OnP2MeetingRecordingReadyV1(handler func(ctx context.Context, event *larkvc.P2MeetingRecordingReadyV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_ready_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_ready_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_ready_v1"] = larkvc.NewP2MeetingRecordingReadyV1Handler(handler)
	return dispatcher
}

// 录制开始
//
// - 发生在开始录制时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/recording_started
func (dispatcher *EventDispatcher) OnP2MeetingRecordingStartedV1(handler func(ctx context.Context, event *larkvc.P2MeetingRecordingStartedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.recording_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.recording_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.recording_started_v1"] = larkvc.NewP2MeetingRecordingStartedV1Handler(handler)
	return dispatcher
}

// 屏幕共享结束
//
// - 发生在屏幕共享结束时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/share_ended
func (dispatcher *EventDispatcher) OnP2MeetingShareEndedV1(handler func(ctx context.Context, event *larkvc.P2MeetingShareEndedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.share_ended_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.share_ended_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.share_ended_v1"] = larkvc.NewP2MeetingShareEndedV1Handler(handler)
	return dispatcher
}

// 屏幕共享开始
//
// - 发生在屏幕共享开始时
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/share_started
func (dispatcher *EventDispatcher) OnP2MeetingShareStartedV1(handler func(ctx context.Context, event *larkvc.P2MeetingShareStartedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.meeting.share_started_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.meeting.share_started_v1")
	}
	dispatcher.eventType2EventHandler["vc.meeting.share_started_v1"] = larkvc.NewP2MeetingShareStartedV1Handler(handler)
	return dispatcher
}

// 创建会议室
//
// - 当创建会议室时，会触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/events/created
func (dispatcher *EventDispatcher) OnP2RoomCreatedV1(handler func(ctx context.Context, event *larkvc.P2RoomCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.room.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.room.created_v1")
	}
	dispatcher.eventType2EventHandler["vc.room.created_v1"] = larkvc.NewP2RoomCreatedV1Handler(handler)
	return dispatcher
}

// 删除会议室
//
// - 当删除会议室时，会触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/events/deleted
func (dispatcher *EventDispatcher) OnP2RoomDeletedV1(handler func(ctx context.Context, event *larkvc.P2RoomDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.room.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.room.deleted_v1")
	}
	dispatcher.eventType2EventHandler["vc.room.deleted_v1"] = larkvc.NewP2RoomDeletedV1Handler(handler)
	return dispatcher
}

// 更新会议室
//
// - 当更新会议室时，会触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/events/updated
func (dispatcher *EventDispatcher) OnP2RoomUpdatedV1(handler func(ctx context.Context, event *larkvc.P2RoomUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["vc.room.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "vc.room.updated_v1")
	}
	dispatcher.eventType2EventHandler["vc.room.updated_v1"] = larkvc.NewP2RoomUpdatedV1Handler(handler)
	return dispatcher
}
