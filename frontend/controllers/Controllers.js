import axios from 'axios'
export function GetCurrentUser () {
  localStorage.getItem('token_vaccinations')
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.get('http://localhost:8084/jwt/encodeJwt', {responeType: 'application/json'})
}
export function addPrivika (vaccination, date, preparat, seria, doza, numberkart) {
  localStorage.getItem('token_vaccinations')
  var dataForm = new FormData()
  dataForm.append('vaccination', vaccination)
  dataForm.append('date', date)
  dataForm.append('preparat', preparat)
  dataForm.append('seria', seria)
  dataForm.append('doza', doza)
  dataForm.append('numberkart', numberkart)
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.post('http://localhost:8084/jwt/addPrivivka', dataForm, {responeType: 'application/json'})
}
export function setItemsDefaultVACCINATIONS (vaccination, numberKart) {
  console.log(vaccination, numberKart)
  localStorage.getItem('token_vaccinations')
  var dataForm = new FormData()
  dataForm.append('type_vaccination', vaccination)
  dataForm.append('numberKart', numberKart)
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
  return axios.post('http://localhost:8084/jwt/getPrivivka', dataForm, {responeType: 'application/json'})
}
