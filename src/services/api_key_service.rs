use ring::digest;
use std::env;

pub fn generate_key() -> String {
    let api_key_secret = env::var("API_KEY_SECRET");
    let digest_result = digest::digest(&digest::SHA256, api_key_secret.unwrap().as_bytes());
    hex::encode(digest_result.as_ref())
}

pub fn validate_key(key: &str) -> bool {
    let api_key_secret = env::var("API_KEY_SECRET");
    let digest_result = digest::digest(&digest::SHA256, api_key_secret.unwrap().as_bytes());
    let expected_key = hex::encode(digest_result.as_ref());
    key == expected_key
}
