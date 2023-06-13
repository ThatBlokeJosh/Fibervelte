<script>
  import Router from "svelte-spa-router";

  import Home from "./routes/Home.svelte";
  import Gen from "./routes/Generate.svelte";

  import NotFound from "./routes/NotFound.svelte";

  let routes = {
    "/": Home,
    "/gen": Gen,

    "*": NotFound

  }
</script>

<main>
  <Router {routes}/>
</main>



