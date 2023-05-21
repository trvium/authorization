use dotenv::dotenv;
use jsonwebtoken::{decode, DecodingKey, Validation};
use serde::{Deserialize, Serialize};
use std::env;

#[derive(Debug, Deserialize, Serialize)]
pub struct Claims {
    pub email: String,
}

pub fn decode_token(token: &str) -> Option<Claims> {
    dotenv().ok();
    let jwt_secret = env::var("JWT_SECRET").unwrap();
    let decoding_key = DecodingKey::from_secret(jwt_secret.as_bytes());

    let validation = Validation::default();

    match decode::<Claims>(token, &decoding_key, &validation) {
        Ok(token_data) => Some(token_data.claims),
        Err(_) => None,
    }
}
