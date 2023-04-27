use crate::services;
use rocket::http::Status;
use rocket::response::status;
use rocket::serde::json::Json;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct ApiKeyResponse {
    api_key: String,
}

#[post("/")]
pub fn generate_key() -> status::Custom<Json<ApiKeyResponse>> {
    let api_key = services::api_key_service::generate_key();
    let response = ApiKeyResponse { api_key };
    status::Custom(Status::Created, Json(response))
}

#[get("/<key>")]
pub fn validate_key(key: String) -> status::Custom<Json<bool>> {
    let is_valid = services::api_key_service::validate_key(&key);
    status::Custom(Status::Ok, Json(is_valid))
}
