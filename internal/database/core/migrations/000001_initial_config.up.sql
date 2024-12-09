BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    token UUID PRIMARY KEY uuid_generate_v4(),            
    email VARCHAR(255) NOT NULL UNIQUE, 
    icon VARCHAR(50) NOT NULL CHECK (icon IN ('CHICKEN', 'BEEF', 'PORK', 'FISH', 'SHRIMP', 'LAMB', 'DUCK', 'CRAB', 'SQUID', 'GOAT')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    last_login TIMESTAMP DEFAULT NULL, 
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE restaurants (
    token UUID PRIMARY KEY uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    slogo TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE dishes (
    token UUID PRIMARY KEY uuid_generate_v4(),
    restaurant_token UUID NOT NULL, 
    category VARCHAR(50) NOT NULL CHECK (category IN ('APPETIZERS', 'MAIN_COURSES', 'SIDE_DISHES', 'DESSERTS', 'BEVERAGES', 'SNACKS', 'VEGETARIAN', 'VEGAN', 'GLUTEN_FREE', 'LOW_CARB', 'BREAKFAST')),
    title VARCHAR(255) NOT NULL,
    quick_description TEXT NOT NULL,
    long_description TEXT,
    restrictions TEXT,
    price VARCHAR(50) NOT NULL,
    image_url TEXT NOT NULL,
    video_url TEXT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT fk_restaurant FOREIGN KEY (restaurant_token) REFERENCES restaurants(token) ON DELETE CASCADE
);

CREATE TABLE comments (
    token UUID PRIMARY KEY uuid_generate_v4(),    
    restaurant_token UUID NOT NULL, 
    dish_token UUID NOT NULL,           
    user_token UUID NOT NULL,          
    message TEXT NOT NULL,             
    note INT NOT NULL CHECK (note >= 0 AND note <= 10), 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    active BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT fk_restaurant FOREIGN KEY (restaurant_token) REFERENCES restaurants(token) ON DELETE CASCADE,
    CONSTRAINT fk_dish FOREIGN KEY (dish_token) REFERENCES dishes(token) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_token) REFERENCES users(token) ON DELETE CASCADE
);

INSERT INTO users (token, email, icon)
VALUES
    (uuid_generate_v4(), 'joao@example.com', 'CHICKEN'),
    (uuid_generate_v4(), 'maria@example.com', 'BEEF');

INSERT INTO restaurants (token, name, slogo) VALUES
    ('3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'gourmet-paradise', 'Gourmet Paradise');

INSERT INTO dishes (token, restaurant_token, category, title, quick_description, long_description, restrictions, price, image_url) VALUES
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'APPETIZERS', 'Bruschetta', 'Deliciosa entrada italiana.', 'Fatias de pão torrado cobertas com tomate, alho e manjericão.', NULL, '12.99', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp'),
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'APPETIZERS', 'Spring Rolls', 'Rolinho primavera fresco.', 'Rolinho de vegetais com molho agridoce.', 'Vegano', '10.50', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp'),
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'MAIN_COURSES', 'Spaghetti Carbonara', 'Clássico prato italiano.', 'Espaguete ao molho cremoso com bacon.', 'Contém glúten e lactose', '25.00', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp'),
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'MAIN_COURSES', 'Steak au Poivre', 'Bife ao molho de pimenta.', 'Filé mignon com molho cremoso de pimenta preta.', 'Contém lactose', '45.90', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp'),
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'DESSERTS', 'Tiramisu', 'Sobremesa clássica italiana.', 'Camadas de biscoito, creme mascarpone e café.', 'Contém cafeína', '15.00', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp'),
    (uuid_generate_v4(), '3fbd6a6f-5336-4dac-a653-88eef1e1d56f', 'DESSERTS', 'Petit Gateau', 'Sobremesa francesa com sorvete.', 'Bolo de chocolate quente com sorvete de creme.', NULL, '18.50', 'https://storage.googleapis.com/restaurant-app/bruschetta.webp');

INSERT INTO comments (token, restaurant_token, dish_token, user_token, message, note)
VALUES
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'joao@example.com' LIMIT 1),
     'Excelente sabor, textura perfeita.', 
     5
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'maria@example.com' LIMIT 1),
     'Muito bom, mas poderia ter menos alho.', 
     4
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'joao@example.com' LIMIT 1),
     'Crocante e fresca, ótima entrada.', 
     5
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'maria@example.com' LIMIT 1),
     'Nada mal, mas já comi melhores.', 
     3
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'joao@example.com' LIMIT 1),
     'A combinação de tomate e manjericão é incrível!', 
     5
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'maria@example.com' LIMIT 1),
     'Boa, porém um pouco salgada.', 
     2
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'joao@example.com' LIMIT 1),
     'Adoro o sabor do alho e do azeite.', 
     4
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'maria@example.com' LIMIT 1),
     'A textura poderia ser mais macia.', 
     3
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'joao@example.com' LIMIT 1),
     'Minha entrada favorita do cardápio!', 
     5
    ),
    (uuid_generate_v4(),
    '3fbd6a6f-5336-4dac-a653-88eef1e1d56f',
     (SELECT token FROM dishes WHERE title = 'Bruschetta' LIMIT 1),
     (SELECT token FROM users WHERE email = 'maria@example.com' LIMIT 1),
     'Bom, mas prefiro outra entrada.', 
     3
    );

COMMIT;
