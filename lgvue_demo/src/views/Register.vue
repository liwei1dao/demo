<template>
  <layout-s>
    <v-card min-width="450px"
            class="float-none">
      <v-toolbar color="primary"
                 flat>
        <v-toolbar-title>{{$t('Register')}}</v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <v-form>
          <v-text-field v-model="phonoremail"
                        :error-messages="phonOrEmailErrors"
                        label="PhonOrEmail"
                        required
                        @input="$v.phonoremail.$touch()"
                        @blur="$v.phonoremail.$touch()"></v-text-field>
          <v-text-field v-model="password"
                        :error-messages="passwordErrors"
                        :counter="10"
                        label="Password"
                        :append-icon="passwordeye ? 'mdi-eye' : 'mdi-eye-off'"
                        :type="passwordeye ? 'text' : 'password'"
                        required
                        @click:append="passwordeye = !passwordeye"
                        @input="$v.password.$touch()"
                        @blur="$v.password.$touch()"></v-text-field>
          <v-text-field v-model="confirmpassword"
                        :error-messages="confirmpasswordErrors"
                        :counter="10"
                        label="ConfirmPassword"
                        :append-icon="passwordeye ? 'mdi-eye' : 'mdi-eye-off'"
                        :type="passwordeye ? 'text' : 'password'"
                        required
                        @click:append="passwordeye = !passwordeye"
                        @input="$v.confirmpassword.$touch()"
                        @blur="$v.confirmpassword.$touch()"></v-text-field>
          <v-text-field v-model="captcha"
                        :error-messages="captchaErrors"
                        :counter="4"
                        label="Captcha"
                        required
                        @input="$v.captcha.$touch()"
                        @blur="$v.captcha.$touch()">
            <template v-slot:append>
              <v-btn color="primary"
                     @click="getcaptcha"
                     dark>{{$t('Captcha')}}</v-btn>
            </template>
          </v-text-field>
          <v-btn class="mr-4"
                 color="warning"
                 @click="registerbycaptcha"
                 block>
            {{$t('Register')}}
          </v-btn>
          <v-btn class="mt-4"
                 color="grey lighten-1"
                 to="/login"
                 block>
            {{$t('Login')}}
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </layout-s>
</template>

<script>
import { required, sameAs, minLength, email, numeric, or } from 'vuelidate/lib/validators'
import { sendemailcaptcha } from '@/api/api'
import { registerbycaptcha } from '@/api/user'
import LayoutS from '@/layout/LayoutS.vue'
export default {
  name: "Register",
  components: {
    LayoutS
  },
  validations: {
    phonoremail: {
      required,
      phonoremail: or(email, numeric),
    },
    password: { required, minLength: minLength(6) },
    confirmpassword: {
      sameAsPassword: sameAs('password')
    },
    captcha: {
      required,
      numeric,
      length: (value) => value.length == 4
    },
  },
  data: () => {
    return {
      phonoremail: "",
      password: "",
      confirmpassword: "",
      captcha: "",
      passwordeye: false,
      captchaDisabled: false,
      captchatimer: null,
      captchaLiftTime: 60,
    }
  },
  computed: {
    phonOrEmailErrors () {
      const errors = []
      if (!this.$v.phonoremail.$dirty) return errors
      !this.$v.phonoremail.required && errors.push(this.$t('ValError_Requiredl'))
      !this.$v.phonoremail.phonoremail && errors.push(this.$t('ValError_PhonOrEmail'))
      return errors
    },
    passwordErrors () {
      const errors = []
      if (!this.$v.password.$dirty) return errors
      !this.$v.password.required && errors.push(this.$t('ValError_Requiredl'))
      !this.$v.password.minLength && errors.push(this.$t('ValError_Password'))
      return errors
    },
    confirmpasswordErrors () {
      const errors = []
      if (!this.$v.confirmpassword.$dirty) return errors
      !this.$v.confirmpassword.sameAsPassword && errors.push(this.$t('ValError_ConfirmPassword'))
      return errors
    },
    captchaErrors () {
      const errors = []
      if (!this.$v.captcha.$dirty) return errors
      !this.$v.captcha.required && errors.push(this.$t('ValError_Requiredl'))
      !this.$v.captcha.numeric && errors.push(this.$t('ValError_CaptchaNumber'))
      !this.$v.captcha.length && errors.push(this.$t('ValError_CaptchaLength'))
      return errors
    },
  },
  methods: {
    getcaptcha () {
      this.$v.phonoremail.$touch()
      if (!this.$v.phonoremail.$invalid) {
        this.captchaDisabled = true
        sendemailcaptcha({ PhonOrEmail: this.phonoremail, CaptchaType: 0 }).then(response => {
          const { message } = response
          this.$message.success(message)
          if (!this.captchatimer) {//启动计时器
            this.captchatimer = setInterval(() => {
              if (this.captchaLiftTime > 0 && this.captchaLiftTime <= 60) {
                this.captchaLiftTime--;
                if (this.captchaLiftTime == 0) {
                  clearInterval(this.captchatimer);
                  this.captchaLiftTime = 60;
                  this.captchatimer = null;
                  this.captchaDisabled = false;
                }
              }
            }, 1000)
          }
        }).catch(error => {
          this.$message.error(error.message)
          this.captchaDisabled = false
          this.captchaLiftTime = 60
        })
      }
    },
    registerbycaptcha () {
      this.$v.$touch()
      if (!this.$v.phonoremail.$invalid &&
        !this.$v.password.$invalid &&
        !this.$v.confirmpassword.$invalid &&
        !this.$v.captcha.$invalid
      ) {
        registerbycaptcha({ PhonOrEmail: this.phonoremail, Password: this.password, Captcha: this.captcha }).then(response => {
          const { message } = response
          this.$message.success(message)
        }).catch(error => {
          this.$message.error(error.message)
        })
      }
    }
  }
}
</script>

<style>
</style>