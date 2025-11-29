CREATE TABLE agent_versions (
                                id VARCHAR(36) PRIMARY KEY NOT NULL COMMENT '版本唯一ID',
                                agent_id VARCHAR(36) NOT NULL COMMENT '关联的Agent ID',
                                name VARCHAR(255) NOT NULL COMMENT 'Agent名称',
                                avatar VARCHAR(255) COMMENT 'Agent头像URL',
                                description TEXT COMMENT 'Agent描述',
                                version_number VARCHAR(20) NOT NULL COMMENT '版本号，如1.0.0',
                                system_prompt TEXT COMMENT 'Agent系统提示词',
                                welcome_message TEXT COMMENT '欢迎消息',
                                tool_ids JSON COMMENT 'Agent可使用的工具ID列表，JSON数组格式',
                                knowledge_base_ids JSON COMMENT '关联的知识库ID列表，JSON数组格式',
                                change_log TEXT COMMENT '版本更新日志',
                                publish_status INT DEFAULT 1 COMMENT '发布状态：1-审核中, 2-已发布, 3-拒绝, 4-已下架',
                                reject_reason TEXT COMMENT '审核拒绝原因',
                                review_time DATETIME COMMENT '审核时间',
                                published_at DATETIME COMMENT '发布时间',
                                user_id VARCHAR(36) NOT NULL COMMENT '创建者用户ID',
                                tool_preset_params JSON COMMENT '工具预设参数',
                                multi_modal TINYINT(1) DEFAULT 0 COMMENT '是否支持多模态',
                                created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                deleted_at DATETIME COMMENT '逻辑删除时间',

                                INDEX idx_agent_versions_agent_id (agent_id),
                                INDEX idx_agent_versions_user_id (user_id),
                                INDEX idx_agent_versions_publish_status (publish_status),
                                INDEX idx_agent_versions_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Agent版本实体类，代表一个Agent的发布版本';


CREATE TABLE agent_workspace (
                                 id VARCHAR(36) PRIMARY KEY NOT NULL COMMENT '主键ID',
                                 agent_id VARCHAR(36) NOT NULL COMMENT 'Agent ID',
                                 user_id VARCHAR(36) NOT NULL COMMENT '用户ID',
                                 llm_model_config JSON COMMENT '模型配置，JSON格式',
                                 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 deleted_at DATETIME COMMENT '逻辑删除时间',

                                 INDEX idx_agent_workspace_agent_id (agent_id),
                                 INDEX idx_agent_workspace_user_id (user_id),
                                 INDEX idx_agent_workspace_agent_user (agent_id, user_id),
                                 INDEX idx_agent_workspace_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Agent工作区实体类，用于记录用户添加到工作区的Agent';

CREATE TABLE agents (
                        id VARCHAR(36) PRIMARY KEY NOT NULL COMMENT 'Agent唯一ID',
                        name VARCHAR(255) NOT NULL COMMENT 'Agent名称',
                        avatar VARCHAR(255) COMMENT 'Agent头像URL',
                        description TEXT COMMENT 'Agent描述',
                        system_prompt TEXT COMMENT 'Agent系统提示词',
                        welcome_message TEXT COMMENT '欢迎消息',
                        tool_ids JSON COMMENT '工具ID列表，JSON数组格式',
                        published_version VARCHAR(36) COMMENT '当前发布的版本ID',
                        enabled TINYINT(1) DEFAULT 1 COMMENT 'Agent状态：1-启用，0-禁用',
                        user_id VARCHAR(36) NOT NULL COMMENT '创建者用户ID',
                        tool_preset_params JSON COMMENT '工具预设参数',
                        multi_modal TINYINT(1) DEFAULT 0 COMMENT '是否支持多模态',
                        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        deleted_at DATETIME COMMENT '逻辑删除时间',
                        knowledge_base_ids JSON COMMENT '关联的知识库ID列表，JSON数组格式，用于RAG功能',

                        INDEX idx_agents_user_id (user_id),
                        INDEX idx_agents_enabled (enabled),
                        INDEX idx_agents_created_at (created_at),
                        INDEX idx_agents_published_version (published_version),
                        INDEX idx_agents_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Agent实体类，代表一个AI助手';