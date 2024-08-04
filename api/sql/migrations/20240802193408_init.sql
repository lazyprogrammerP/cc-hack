-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    department (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL);

CREATE TABLE
    doctor (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, background TEXT NOT NULL, education_details TEXT NOT NULL, experience SMALLINT NOT NULL, fees INT NOT NULL, department_id INT NOT NULL REFERENCES department (id) ON DELETE CASCADE ON UPDATE CASCADE);

CREATE TABLE
    amenity (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL);

CREATE TABLE
    insurance_partner (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL);

CREATE TABLE
    hospital (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL, lat DECIMAL(10, 8) NOT NULL, long DECIMAL(11, 8) NOT NULL, cost INT NOT NULL, accreditation VARCHAR(255), wait_time INT NOT NULL, capacity INT NOT NULL, site_url VARCHAR(255), contact_number VARCHAR(15) NOT NULL, ratings DECIMAL(3, 2) NOT NULL);

CREATE TABLE
    address (id SERIAL PRIMARY KEY, street VARCHAR(255) NOT NULL, landmark VARCHAR(255), city VARCHAR(100) NOT NULL, pincode VARCHAR(10) NOT NULL, hospital_id INT NOT NULL UNIQUE REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE);

CREATE TABLE
    hospital_image (id SERIAL PRIMARY KEY, src VARCHAR(255) NOT NULL, hospital_id INT NOT NULL REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE);

CREATE TABLE
    hospital_departments (id SERIAL PRIMARY KEY, department_id INT NOT NULL REFERENCES department (id) ON DELETE CASCADE ON UPDATE CASCADE, hospital_id INT NOT NULL REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE, UNIQUE (department_id, hospital_id));

CREATE TABLE
    hospital_doctors (id SERIAL PRIMARY KEY, doctor_id INT NOT NULL REFERENCES doctor (id) ON DELETE CASCADE ON UPDATE CASCADE, hospital_id INT NOT NULL REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE, UNIQUE (doctor_id, hospital_id));

CREATE TABLE
    hospital_amenities (id SERIAL PRIMARY KEY, amenity_id INT NOT NULL REFERENCES amenity (id) ON DELETE CASCADE ON UPDATE CASCADE, hospital_id INT NOT NULL REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE, UNIQUE (amenity_id, hospital_id));

CREATE TABLE
    hospital_insurance_partners (id SERIAL PRIMARY KEY, insurance_partner_id INT NOT NULL REFERENCES insurance_partner (id) ON DELETE CASCADE ON UPDATE CASCADE, hospital_id INT NOT NULL REFERENCES hospital (id) ON DELETE CASCADE ON UPDATE CASCADE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hospital_insurance_partners;

DROP TABLE IF EXISTS hospital_amenities;

DROP TABLE IF EXISTS hospital_doctors;

DROP TABLE IF EXISTS hospital_departments;

DROP TABLE IF EXISTS insurance_partner;

DROP TABLE IF EXISTS amenity;

DROP TABLE IF EXISTS doctor;

DROP TABLE IF EXISTS department;

DROP TABLE IF EXISTS address;

DROP TABLE IF EXISTS hospital_image;

DROP TABLE IF EXISTS hospital;

-- +goose StatementEnd