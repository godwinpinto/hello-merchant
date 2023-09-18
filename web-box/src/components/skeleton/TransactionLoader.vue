<script setup lang="ts">
const arraySkeleton = [0, 0, 0, 0];
</script>
<template>
    <div class="min-h-72 h-80 overflow-y-scroll pr-4" id="about">
        <ul role="list" class="divide-y divide-gray-200 dark:divide-gray-700 animate-pulse">
            <li class="py-3 sm:py-4" v-for="(val, index) in arraySkeleton">
                <div class="flex items-center space-x-4">
                    <div
                        class="relative inline-flex items-center justify-center w-10 h-10 overflow-hidden bg-gray-100 rounded-full dark:bg-gray-600">
                        <span class="font-medium text-gray-600 dark:text-gray-300"></span>
                    </div>
                    <div class="flex-1 min-w-10">
                        <div class="bg-gray-200 rounded-full dark:bg-gray-700  h-4 w-40 mb-2">
                        </div>
                        <div class="bg-gray-200 rounded-full dark:bg-gray-700  h-4 w-72">
                        </div>
                    </div>
                    <div class="inline-flex items-center bg-gray-200 rounded-full dark:bg-gray-700 h-8 w-20">
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>