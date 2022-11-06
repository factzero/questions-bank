import useUserStore from "./user";
import useRouterStore from "./router";

const useStore = () => ({
  userStore: useUserStore(),
  routerStore: useRouterStore(),
});

export default useStore;
