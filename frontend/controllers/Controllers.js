import axios from 'axios'
export function GetCurrentUser () {
  localStorage.getItem('token_vaccinations')
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
}
export function addPrivika (vaccination, data, preparat, seria, doza, numberkart) {
  localStorage.getItem('token_vaccinations')
  var dataForm = new FormData()
  dataForm.append('vaccination', vaccination)
  dataForm.append('data', data)
  dataForm.append('preparat', preparat)
  dataForm.append('seria', seria)
  dataForm.append('doza', doza)
  dataForm.append('numberkart', numberkart)
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.post('http://localhost:8084/jwt/encodeJwt', dataForm, {responeType: 'application/json'})
}
export function setItemsDefaultVACCINATIONS (idx) {
  localStorage.getItem('token_vaccinations')
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  switch (idx) {
    case 0: return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
    case 1: return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
    case 2: return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
    case 3: return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
  }

}
