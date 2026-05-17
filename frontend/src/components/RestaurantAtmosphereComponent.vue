<template>
  <section class="atmosphere">
    <h1>Атмосфера</h1>

    <div class="atmosphere-gallery">
      <img src="../assets/olio_atm_1.png" alt="" />
      <img src="../assets/olio_atm_2.png" alt="" />
      <img src="../assets/olio_atm_3.png" alt="" />
    </div>
  </section>
</template>

<style>
.atmosphere {
  height: 350px;
  text-align: center;
  margin-bottom: 40px;
}

.atmosphere-gallery {
  display: flex;
  justify-content: center;
  gap: 50px;
  margin-top: 30px;
}

.atmosphere-gallery img {
  width: 250px;
}
</style>
