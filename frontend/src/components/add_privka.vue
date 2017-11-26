<btn-add-privivka>

</btn-add-privivka>

<script>
  import Vue from 'vue'
  import 'bootstrap/dist/css/bootstrap.css'
  import 'bootstrap-vue/dist/bootstrap-vue.css'
  import * as controllers from '../../controllers/Controllers.js'

  Vue.component('btn-add-privivka', {
    template: '#add-privivka',
    data: function () {
      return {
        currentPrivivka: null,
        selectedDate: null,
        selectedPreparat: null,
        selectedSeria: null,
        selectedDoza: null
      }
    },
    methods: {
      addPrivivka: function () {
        var ss = this
        controllers.addPrivika(this.currentPrivivka,
          this.selectedDate,
          this.selectedPreparat,
          this.selectedSeria,
          this.selectedDoza,
          this.NumberKart
        ).then(function (response) {
          console.log(response.status)
          if (response.status === 200) {
            ss.setItemsDefaultVACCINATIONS('АДСМ')
          }
          if (response.status === 404) {

          }
        }).catch(function (error) {
          console.log(error.response)
          alert('Произошла ошибка, свяжитесь с системным администратором ' + error)
        })
      }
    }
  })
</script>
