-- 创建表单定义表
CREATE TABLE forms (
  id BIGSERIAL PRIMARY KEY,
  tenant_id BIGINT NOT NULL,
  form_key VARCHAR(64) UNIQUE NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  category_id BIGINT,
  layout TEXT DEFAULT 'single-column',
  status SMALLINT DEFAULT 1 CHECK (status IN (0, 1)),
  version INT DEFAULT 1,
  deleted_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);

-- 创建字段定义表
CREATE TABLE form_fields (
  id BIGSERIAL PRIMARY KEY,
  form_id BIGINT NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
  group_key TEXT,
  group_label TEXT,
  field_key TEXT NOT NULL,
  label TEXT NOT NULL,
  type TEXT NOT NULL,
  required BOOLEAN DEFAULT FALSE,
  options JSONB,
  default_value JSONB,
  placeholder TEXT,
  validation JSONB,
  visible_condition JSONB,
  editable_condition JSONB,
  order_num INT DEFAULT 0,
  encrypted BOOLEAN DEFAULT FALSE,
  deleted_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now(),
  UNIQUE(form_id, field_key)
);

-- 多语言字段支持
CREATE TABLE form_field_i18n (
  id BIGSERIAL PRIMARY KEY,
  field_id BIGINT NOT NULL REFERENCES form_fields(id) ON DELETE CASCADE,
  lang_code VARCHAR(8) NOT NULL,
  label TEXT,
  placeholder TEXT,
  UNIQUE(field_id, lang_code)
);

-- 表单提交记录
CREATE TABLE form_entries (
  id BIGSERIAL PRIMARY KEY,
  form_id BIGINT NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
  tenant_id BIGINT NOT NULL,
  user_id BIGINT,
  data JSONB NOT NULL,
  status TEXT DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'approved', 'rejected')),
  version INT DEFAULT 1,
  deleted_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);

-- 审批流程表
CREATE TABLE form_entry_audit (
  id BIGSERIAL PRIMARY KEY,
  entry_id BIGINT NOT NULL REFERENCES form_entries(id) ON DELETE CASCADE,
  step INT NOT NULL,
  approver_id BIGINT NOT NULL,
  status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
  comment TEXT,
  approved_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 表单结构快照（历史版本）
CREATE TABLE form_versions (
  id BIGSERIAL PRIMARY KEY,
  form_id BIGINT NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
  version INT NOT NULL,
  schema_snapshot JSONB NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 表单提交历史记录
CREATE TABLE form_entry_history (
  id BIGSERIAL PRIMARY KEY,
  entry_id BIGINT NOT NULL REFERENCES form_entries(id) ON DELETE CASCADE,
  form_id BIGINT NOT NULL,
  changed_by BIGINT,
  old_data JSONB,
  new_data JSONB,
  change_type TEXT CHECK (change_type IN ('create', 'update', 'delete')),
  created_at TIMESTAMPTZ DEFAULT now()
);

-- 索引优化
CREATE INDEX idx_form_entries_data_gin ON form_entries USING GIN (data);
CREATE INDEX idx_form_entries_status ON form_entries(status);
CREATE INDEX idx_form_entries_form_user ON form_entries(form_id, user_id);
CREATE INDEX idx_form_fields_form_order ON form_fields(form_id, order_num);
CREATE INDEX idx_form_entry_history_entry ON form_entry_history(entry_id);
CREATE INDEX idx_form_versions_form_version ON form_versions(form_id, version DESC);