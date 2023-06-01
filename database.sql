create table group_member_info
(
    GroupId            int                          null comment '群号',
    MemberId           int                          null comment '成员账号',
    MemberName         varchar(256) charset utf8mb4 null comment '群名片',
    SpecialTitle       varchar(256) charset utf8mb4 null comment '特殊称号',
    JoinTimestamp      timestamp                    null comment '加群时间戳',
    LastSpeakTimestamp timestamp                    null comment '最后发言时间',
    MuteTimeRemaining  int                          null comment '剩余禁言时长'
);

create table `groups`
(
    id         int          null comment '群号',
    name       varchar(255) null comment '群名称',
    permission varchar(255) null comment 'bot在群聊中的权限'
);
