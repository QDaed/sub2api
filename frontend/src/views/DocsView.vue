<template>
  <div
    class="relative flex min-h-screen flex-col overflow-hidden bg-gradient-to-br from-gray-50 via-primary-50/30 to-gray-100 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950"
  >
    <!-- Background Decorations -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div
        class="absolute -right-40 -top-40 h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"
      ></div>
      <div
        class="absolute -bottom-40 -left-40 h-96 w-96 rounded-full bg-primary-500/15 blur-3xl"
      ></div>
      <div
        class="absolute left-1/3 top-1/4 h-72 w-72 rounded-full bg-primary-300/10 blur-3xl"
      ></div>
      <div
        class="absolute bottom-1/4 right-1/4 h-64 w-64 rounded-full bg-primary-400/10 blur-3xl"
      ></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(20,184,166,0.03)_1px,transparent_1px),linear-gradient(90deg,rgba(20,184,166,0.03)_1px,transparent_1px)] bg-[size:64px_64px]"
      ></div>
    </div>

    <PublicNavbar
      :site-name="siteName"
      :site-logo="siteLogo"
      :is-dark="isDark"
      :is-authenticated="isAuthenticated"
      :dashboard-path="dashboardPath"
      :user-initial="userInitial"
      @toggle-theme="toggleTheme"
    />

    <main class="relative z-10 flex-1 px-6 py-16">
      <div class="mx-auto max-w-6xl">
        <section
          class="mb-12 flex flex-col items-center justify-between gap-12 lg:flex-row lg:gap-16"
        >
          <div class="flex-1 text-center lg:text-left">
            <p
              class="mb-4 inline-flex items-center gap-2 rounded-full border border-primary-200 bg-white/70 px-4 py-2 text-sm font-medium text-primary-700 shadow-sm backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/70 dark:text-primary-300"
            >
              <Icon name="book" size="sm" />
              {{ t("docs.badge") }}
            </p>
            <h1
              class="mb-4 text-4xl font-bold text-gray-900 dark:text-white md:text-5xl lg:text-6xl"
            >
              {{ t("docs.title") }}
            </h1>
            <p class="mb-8 text-lg text-gray-600 dark:text-dark-300 md:text-xl">
              {{ t("docs.subtitle") }}
            </p>

            <div
              class="flex flex-col items-center gap-3 sm:flex-row lg:justify-start"
            >
              <a
                v-if="docUrl"
                :href="docUrl"
                target="_blank"
                rel="noopener noreferrer"
                class="btn btn-primary px-8 py-3 text-base shadow-lg shadow-primary-500/30"
              >
                {{ t("docs.openExternal") }}
                <Icon
                  name="externalLink"
                  size="md"
                  class="ml-2"
                  :stroke-width="2"
                />
              </a>
              <router-link
                :to="isAuthenticated ? dashboardPath : '/login'"
                class="inline-flex items-center rounded-full bg-gray-900 px-6 py-3 text-sm font-medium text-white transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
              >
                {{
                  isAuthenticated
                    ? t("home.goToDashboard")
                    : t("home.getStarted")
                }}
              </router-link>
            </div>
          </div>

          <div class="flex flex-1 justify-center lg:justify-end">
            <div
              class="w-full max-w-md rounded-2xl border border-gray-200/50 bg-white/70 p-6 shadow-xl shadow-primary-500/10 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/70"
            >
              <div class="mb-5 flex items-center gap-3">
                <div
                  class="flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-primary-500 to-primary-600 text-white shadow-lg shadow-primary-500/30"
                >
                  <Icon name="document" size="lg" :stroke-width="1.8" />
                </div>
                <div class="text-left">
                  <h2
                    class="text-lg font-semibold text-gray-900 dark:text-white"
                  >
                    {{ t("docs.quickStart.title") }}
                  </h2>
                  <p class="text-sm text-gray-500 dark:text-dark-400">
                    {{ t("docs.quickStart.description") }}
                  </p>
                </div>
              </div>

              <div class="space-y-3">
                <div
                  v-for="(step, index) in quickStartSteps"
                  :key="step.title"
                  class="flex gap-3 rounded-xl border border-gray-200/50 bg-white/70 p-4 text-left dark:border-dark-700/50 dark:bg-dark-900/50"
                >
                  <span
                    class="flex h-7 w-7 shrink-0 items-center justify-center rounded-full bg-primary-100 text-xs font-semibold text-primary-700 dark:bg-primary-900/30 dark:text-primary-300"
                  >
                    {{ index + 1 }}
                  </span>
                  <div>
                    <h3
                      class="text-sm font-semibold text-gray-900 dark:text-white"
                    >
                      {{ step.title }}
                    </h3>
                    <p class="mt-1 text-sm text-gray-600 dark:text-dark-400">
                      {{ step.description }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section
          v-if="docUrl"
          class="mb-12 overflow-hidden rounded-2xl border border-gray-200/50 bg-white/70 shadow-xl shadow-primary-500/10 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/70"
        >
          <div
            class="flex items-center justify-between border-b border-gray-200/50 px-5 py-4 dark:border-dark-700/50"
          >
            <div>
              <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
                {{ t("docs.embedded.title") }}
              </h2>
              <p class="text-sm text-gray-500 dark:text-dark-400">
                {{ t("docs.embedded.description") }}
              </p>
            </div>
          </div>
          <iframe
            :src="docUrl"
            class="h-[70vh] w-full border-0 bg-white"
            loading="lazy"
            :title="t('docs.title')"
          ></iframe>
        </section>

        <section class="grid gap-5 md:grid-cols-3">
          <article
            v-for="card in docCards"
            :key="card.title"
            class="group rounded-2xl border border-gray-200/50 bg-white/60 p-6 backdrop-blur-sm transition-all duration-300 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/50 dark:bg-dark-800/60"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-primary-500 to-primary-600 text-white shadow-lg shadow-primary-500/30 transition-transform group-hover:scale-110"
            >
              <Icon :name="card.icon" size="md" :stroke-width="1.8" />
            </div>
            <h2
              class="mb-2 text-lg font-semibold text-gray-900 dark:text-white"
            >
              {{ card.title }}
            </h2>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ card.description }}
            </p>
          </article>
        </section>
      </div>
    </main>

    <footer
      class="relative z-10 border-t border-gray-200/50 px-6 py-8 dark:border-dark-800/50"
    >
      <div
        class="mx-auto flex max-w-6xl flex-col items-center justify-center gap-4 text-center sm:flex-row sm:text-left"
      >
        <p class="text-sm text-gray-500 dark:text-dark-400">
          &copy; {{ currentYear }} {{ siteName }}.
          {{ t("home.footer.allRightsReserved") }}
        </p>
        <router-link
          to="/home"
          class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
        >
          {{ t("docs.backHome") }}
        </router-link>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useAuthStore, useAppStore } from "@/stores";
import Icon from "@/components/icons/Icon.vue";
import PublicNavbar from "@/components/layout/PublicNavbar.vue";

type IconName = InstanceType<typeof Icon>["$props"]["name"];

const { t } = useI18n();

const authStore = useAuthStore();
const appStore = useAppStore();

const siteName = computed(
  () =>
    appStore.cachedPublicSettings?.site_name || appStore.siteName || "Sub2API",
);
const siteLogo = computed(
  () => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || "",
);
const docUrl = computed(
  () => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || "",
);

const isDark = ref(document.documentElement.classList.contains("dark"));
const isAuthenticated = computed(() => authStore.isAuthenticated);
const isAdmin = computed(() => authStore.isAdmin);
const dashboardPath = computed(() =>
  isAdmin.value ? "/admin/dashboard" : "/dashboard",
);
const userInitial = computed(() => {
  const user = authStore.user;
  if (!user || !user.email) return "";
  return user.email.charAt(0).toUpperCase();
});
const currentYear = computed(() => new Date().getFullYear());

const quickStartSteps = computed(() => [
  {
    title: t("docs.quickStart.createKey.title"),
    description: t("docs.quickStart.createKey.description"),
  },
  {
    title: t("docs.quickStart.configureClient.title"),
    description: t("docs.quickStart.configureClient.description"),
  },
  {
    title: t("docs.quickStart.monitorUsage.title"),
    description: t("docs.quickStart.monitorUsage.description"),
  },
]);

const docCards = computed<
  Array<{ icon: IconName; title: string; description: string }>
>(() => [
  {
    icon: "key",
    title: t("docs.cards.apiKeys.title"),
    description: t("docs.cards.apiKeys.description"),
  },
  {
    icon: "swap",
    title: t("docs.cards.routing.title"),
    description: t("docs.cards.routing.description"),
  },
  {
    icon: "chart",
    title: t("docs.cards.billing.title"),
    description: t("docs.cards.billing.description"),
  },
]);

function toggleTheme() {
  isDark.value = !isDark.value;
  document.documentElement.classList.toggle("dark", isDark.value);
  localStorage.setItem("theme", isDark.value ? "dark" : "light");
}

function initTheme() {
  const savedTheme = localStorage.getItem("theme");
  if (
    savedTheme === "dark" ||
    (!savedTheme && window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    isDark.value = true;
    document.documentElement.classList.add("dark");
  }
}

onMounted(() => {
  initTheme();
  authStore.checkAuth();

  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings();
  }
});
</script>
