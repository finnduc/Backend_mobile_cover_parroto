-- Roles
INSERT INTO roles (name) VALUES ('admin'), ('user') ON CONFLICT DO NOTHING;

-- Permissions
INSERT INTO permissions (name) VALUES 
('lesson:create'), ('lesson:update'), ('lesson:delete'),
('user:manage'), ('analytics:view')
ON CONFLICT DO NOTHING;

-- Map Admin permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p WHERE r.name = 'admin'
ON CONFLICT DO NOTHING;

-- Basic Categories
INSERT INTO categories (name) VALUES 
('Movies'), ('Daily Conversations'), ('Business English'), ('Music'), ('Science & Tech')
ON CONFLICT DO NOTHING;
