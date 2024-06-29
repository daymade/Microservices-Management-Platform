-- 插入测试数据
INSERT INTO users (username, email, created_at) VALUES
                                                    ('john_doe', 'john.doe@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('jane_smith', 'jane.smith@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('mike_johnson', 'mike.johnson@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('emily_brown', 'emily.brown@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('david_wilson', 'david.wilson@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('sarah_lee', 'sarah.lee@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('chris_taylor', 'chris.taylor@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('lisa_anderson', 'lisa.anderson@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('robert_martinez', 'robert.martinez@example.com', NOW() - INTERVAL '1 year' * random()),
                                                    ('amy_garcia', 'amy.garcia@example.com', NOW() - INTERVAL '1 year' * random());

INSERT INTO services (name, description, owner_id, created_at, updated_at) VALUES
                                                                               ('User Authentication API', 'Secure user authentication and authorization service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Payment Gateway', 'Integrated payment processing solution', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Inventory Management System', 'Real-time inventory tracking and management', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Customer Relationship Management', 'Comprehensive CRM platform', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Data Analytics Engine', 'Advanced data processing and analytics service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Content Delivery Network', 'High-performance content distribution service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Email Marketing Platform', 'Automated email campaign management system', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Machine Learning API', 'Scalable machine learning model deployment service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Social Media Integration', 'Unified social media management and analytics', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Cloud Storage Solution', 'Secure and scalable cloud storage service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Real-time Messaging API', 'High-performance real-time messaging system', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Geolocation Services', 'Accurate location-based services and mapping', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Image Recognition API', 'Advanced image analysis and recognition service', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('Blockchain Integration', 'Secure blockchain-based transaction processing', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random()),
                                                                               ('IoT Device Management', 'Comprehensive IoT device control and monitoring', (SELECT id FROM users ORDER BY random() LIMIT 1), NOW() - INTERVAL '6 months' * random(), NOW() - INTERVAL '1 month' * random());

-- 为每个服务添加1到5个随机版本
DO $$
    DECLARE
        service_id BIGINT;
        version_count INT;
    BEGIN
        FOR service_id IN SELECT id FROM services LOOP
                version_count := floor(random() * 5 + 1);
                FOR i IN 1..version_count LOOP
                        INSERT INTO versions (service_id, number, description, created_at)
                        VALUES (
                                   service_id,
                                   'v' || i || '.' || floor(random() * 10)::TEXT,
                                   'Version ' || i || ' of the service',
                                   NOW() - INTERVAL '3 months' * random()
                               );
                    END LOOP;
            END LOOP;
    END $$;
