<template>
  <v-container class="pa-0 d-flex justify-center fill-height"
               fluid>
    <v-card>
      <v-card-title>
        Message 演示
      </v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        <v-btn class="ma-2"
               color="success"
               @click="showmessage_succ">
          Success
        </v-btn>
        <v-btn class="ma-2"
               color="alert"
               @click="showmessage_alert">
          Alert
        </v-btn>
        <v-btn class="ma-2"
               color="error"
               @click="showmessage_error">
          Error
        </v-btn>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "Message",
  methods: {
    showmessage_succ () {
      this.$message.success("Success")
    },
    showmessage_alert () {
      this.$message.alert("Alert")
    },
    showmessage_error () {
      this.$message.error("Error")
    },
  }
}
</script>

<style>
</style>