import axios from 'axios'
export function GetCurrentUser () {
  localStorage.getItem('token_vaccinations')
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
}
