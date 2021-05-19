<template>
  <layout-s>
    <template v-slot:appbarbutts>
      <select-languge></select-languge>
    </template>
    <template v-slot>
      <v-card min-width="450px"
              class="float-none">
        <v-toolbar color="primary"
                   flat>
          <v-toolbar-title>{{$t('Login')}}</v-toolbar-title>
        </v-toolbar>
        <v-card-text>
          <form>
            <v-text-field v-model="phonoremail"
                          :error-messages="phonoremailErrors"
                          label="PhonOrEmail"
                          required
                          @input="$v.phonoremail.$touch()"
                          @blur="$v.phonoremail.$touch()"></v-text-field>
            <v-text-field v-model="password"
                          :error-messages="passwordErrors"
                          :counter="10"
                          label="Password"
                          required
                          @input="$v.password.$touch()"
                          @blur="$v.password.$touch()"></v-text-field>
            <v-btn class="mr-4"
                   color="warning"
                   @click="loginbypassword"
                   block>
              {{$t('Login')}}
            </v-btn>
            <v-btn class="mt-4"
                   color="grey lighten-1"
                   to="/register"
                   block>
              {{$t('Register')}}
            </v-btn>
          </form>
        </v-card-text>
      </v-card>
    </template>

  </layout-s>
</template>

<script>
import { required, maxLength, email, numeric, or } from 'vuelidate/lib/validators'
import LayoutS from '@/layout/LayoutS.vue'
import SelectLanguge from '@/components/selectlanguge/'
export default {
  name: "Login",
  components: {
    LayoutS,
    SelectLanguge
  },
  validations: {
    password: { required, maxLength: maxLength(10) },
    phonoremail: {
      required,
      phonoremail: or(email, numeric),
    },
    select: { required },
  },
  data: () => {
    return {
      phonoremail: '',
      password: '',
    }
  },
  computed: {
    phonoremailErrors () {
      const errors = []
      if (!this.$v.phonoremail.$dirty) return errors
      !this.$v.phonoremail.required && errors.push('E-mail is required')
      !this.$v.phonoremail.phonoremail && errors.push('Must be valid e-mail')
      return errors
    },
    passwordErrors () {
      const errors = []
      if (!this.$v.password.$dirty) return errors
      !this.$v.password.maxLength && errors.push('Name must be at most 10 characters long')
      !this.$v.password.required && errors.push('Name is required.')
      return errors
    },
  },
  methods: {
    submit () {
      this.$v.$touch()
      if (this.$v.phonoremail.$invalid) {
        console.log("登录请求");
      }
    },
    clear () {
      this.$v.$reset()
      this.name = ''
      this.email = ''
      this.select = null
      this.checkbox = false
    },
    loginbypassword () {
      this.$v.$touch()
      if (!this.$v.phonoremail.$invalid && !this.$v.password.$invalid) {
        this.loading = true
        this.$store.dispatch('user/loginbypassword', { PhonOrEmail: this.phonoremail, Password: this.password })
          .then((data) => {
            this.$store.dispatch('app/setrouters', data.UserRole ? data.UserRole : 0)
            this.$router.addRoutes(this.$store.getters.routers)
            this.loading = false
            this.$router.push({ path: '/index' })
          })
          .catch((error) => {
            this.loading = false
            this.$message.error(error.message)
          }
          )
      }
    }
  }
}
</script>

<style>
</style>