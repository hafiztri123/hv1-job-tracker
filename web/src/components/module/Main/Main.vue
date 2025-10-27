<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useToast } from 'vue-toastification'
import ApplicationServices from '@/services/application.service'
import type { Application, CreateApplicationDto, UpdateApplicationDto } from '@/services/dto/application.dto'
import Button from '@/components/common/Button.vue'
import Input from '@/components/common/Input.vue'
import Form from '@/components/common/Form.vue'
import { Delete24Filled, Edit24Filled, MoreVertical24Filled } from '@vicons/fluent'

interface ErrorResponse {
  response?: {
    data?: {
      message?: string
    }
  }
}

const toast = useToast()

const applications = ref<Application[]>([])
const statusOptions = ref<string[]>([])
const selectedStatus = ref<string>('')
const loading = ref(false)
const showModal = ref(false)
const showDeleteModal = ref(false)
const showBatchDeleteModal = ref(false)
const editingId = ref<string | null>(null)
const deletingId = ref<string | null>(null)
const editingStatusId = ref<string | null>(null)
const openActionDropdown = ref<string | null>(null)
const selectedIds = ref<Set<string>>(new Set())
const showBatchStatusDropdown = ref(false)

const currentPage = ref(1)
const pageSize = ref(10)
const totalCount = ref(0)

const formValue = reactive<CreateApplicationDto>({
  companyName: '',
  positionTitle: '',
  jobUrl: '',
  salaryRange: '',
  location: '',
  status: 'Wishlist',
  notes: '',
  appliedDate: '',
})

const formValidation = ref<Record<string, boolean>>({
  companyName: false,
  positionTitle: false,
})

const canSubmit = computed(() => {
  return formValidation.value.companyName && formValidation.value.positionTitle
})

const totalPages = computed(() => {
  return Math.ceil(totalCount.value / pageSize.value)
})

const offset = computed(() => {
  return (currentPage.value - 1) * pageSize.value
})

const isAllSelected = computed(() => {
  return applications.value.length > 0 && applications.value.every(app => selectedIds.value.has(app.id))
})

const hasSelection = computed(() => {
  return selectedIds.value.size > 0
})

const loadApplications = async () => {
  try {
    loading.value = true
    const response = await ApplicationServices.getApplications(
      selectedStatus.value || undefined,
      pageSize.value,
      offset.value
    )
    applications.value = response.data.data.data || []
    totalCount.value = response.data.data.dataCount || 0
    closeActionDropdown()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to load applications')
  } finally {
    loading.value = false
  }
}

const loadStatusOptions = async () => {
  try {
    const response = await ApplicationServices.getApplicationOptions()
    statusOptions.value = response.data.data.statusOption || []
  } catch (error) {
    console.error('Failed to load status options', error)
  }
}

const resetForm = () => {
  formValue.companyName = ''
  formValue.positionTitle = ''
  formValue.jobUrl = ''
  formValue.salaryRange = ''
  formValue.location = ''
  formValue.status = 'Wishlist'
  formValue.notes = ''
  formValue.appliedDate = ''
  editingId.value = null
}

const openNewForm = () => {
  resetForm()
  showModal.value = true
}

const openEditForm = (app: Application) => {
  editingId.value = app.id
  formValue.companyName = app.companyName
  formValue.positionTitle = app.positionTitle
  formValue.jobUrl = app.jobUrl || ''
  formValue.salaryRange = app.salaryRange || ''
  formValue.location = app.location || ''
  formValue.status = app.status || ''
  formValue.notes = app.notes || ''
  formValue.appliedDate = app.appliedDate || ''
  showModal.value = true
}

const handleSubmit = async () => {
  if (!canSubmit.value) return

  try {
    loading.value = true
    if (editingId.value) {
      const updateDto: UpdateApplicationDto = {
        companyName: formValue.companyName || undefined,
        positionTitle: formValue.positionTitle || undefined,
        jobUrl: formValue.jobUrl || undefined,
        salaryRange: formValue.salaryRange || undefined,
        location: formValue.location || undefined,
        status: formValue.status || undefined,
        notes: formValue.notes || undefined,
        appliedDate: formValue.appliedDate ? new Date(formValue.appliedDate).toISOString() : undefined,
      }
      await ApplicationServices.updateApplication(editingId.value, updateDto)
      toast.success('Application updated successfully')
    } else {
      const createDto: CreateApplicationDto = {
        companyName: formValue.companyName,
        positionTitle: formValue.positionTitle,
        jobUrl: formValue.jobUrl || undefined,
        salaryRange: formValue.salaryRange || undefined,
        location: formValue.location || undefined,
        status: formValue.status || undefined,
        notes: formValue.notes || undefined,
        appliedDate: formValue.appliedDate ? new Date(formValue.appliedDate).toISOString() : undefined,
      }
      await ApplicationServices.createApplication(createDto)
      toast.success('Application created successfully')
    }
    showModal.value = false
    await loadApplications()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to save application')
  } finally {
    loading.value = false
  }
}

const handleDelete = async (id: string) => {
  deletingId.value = id
  showDeleteModal.value = true
}

const confirmDelete = async () => {
  if (!deletingId.value) return

  try {
    loading.value = true
    await ApplicationServices.deleteApplication(deletingId.value)
    toast.success('Application deleted successfully')
    showDeleteModal.value = false
    deletingId.value = null
    if (applications.value.length === 1 && currentPage.value > 1) {
      currentPage.value--
    }
    await loadApplications()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to delete application')
  } finally {
    loading.value = false
  }
}

const cancelDelete = () => {
  showDeleteModal.value = false
  deletingId.value = null
}

const handleStatusChange = async (id: string, newStatus: string) => {
  try {
    loading.value = true
    const updateDto: UpdateApplicationDto = {
      status: newStatus
    }
    await ApplicationServices.updateApplication(id, updateDto)
    toast.success('Status updated successfully')
    await loadApplications()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to update status')
    editingStatusId.value = null
  } finally {
    loading.value = false
  }
}

const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    loadApplications()
  }
}

const changePageSize = async (newSize: number) => {
  pageSize.value = newSize
  currentPage.value = 1
  await loadApplications()
}

const handleStatusFilterChange = (status: string) => {
  selectedStatus.value = status
  currentPage.value = 1
  loadApplications()
}

const toggleActionDropdown = (id: string) => {
  openActionDropdown.value = openActionDropdown.value === id ? null : id
}

const closeActionDropdown = () => {
  openActionDropdown.value = null
}

const registerValidation = (field: string, isValid: boolean) => {
  formValidation.value[field] = isValid
}

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedIds.value.clear()
  } else {
    applications.value.forEach(app => selectedIds.value.add(app.id))
  }
}

const toggleSelectId = (id: string) => {
  if (selectedIds.value.has(id)) {
    selectedIds.value.delete(id)
  } else {
    selectedIds.value.add(id)
  }
}

const clearSelection = () => {
  selectedIds.value.clear()
}

const handleBatchDelete = () => {
  showBatchDeleteModal.value = true
}

const confirmBatchDelete = async () => {
  if (selectedIds.value.size === 0) return

  try {
    loading.value = true
    const ids = Array.from(selectedIds.value)
    await ApplicationServices.batchDeleteApplications(ids)
    toast.success(`${ids.length} application(s) deleted successfully`)
    showBatchDeleteModal.value = false
    clearSelection()
    if (applications.value.length === ids.length && currentPage.value > 1) {
      currentPage.value--
    }
    await loadApplications()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to delete applications')
  } finally {
    loading.value = false
  }
}

const cancelBatchDelete = () => {
  showBatchDeleteModal.value = false
}

const handleBatchStatusChange = async (status: string) => {
  if (selectedIds.value.size === 0) return

  try {
    loading.value = true
    const ids = Array.from(selectedIds.value)
    await ApplicationServices.batchUpdateStatusApplications(ids, status)
    toast.success(`${ids.length} application(s) status updated successfully`)
    showBatchStatusDropdown.value = false
    clearSelection()
    await loadApplications()
  } catch (error) {
    const err = error as ErrorResponse
    toast.error(err.response?.data?.message || 'Failed to update status')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadApplications()
  loadStatusOptions()
})
</script>

<template>
  <div class="h-full flex flex-col overflow-hidden bg-white">
    <div class="flex-1 flex flex-col overflow-hidden">
      <div class="px-6 py-4 border-b">
        <div class="flex justify-between items-center mb-4">
          <h1 class="text-2xl font-bold">Applications</h1>
          <button
            @click="openNewForm"
            class="px-4 py-2 bg-black text-white rounded-md font-semibold hover:bg-gray-800 transition-colors"
          >
            + Add Application
          </button>
        </div>

        <div class="flex gap-2 flex-wrap">
          <button
            @click="handleStatusFilterChange('')"
            :class="[
              'px-3 py-1 rounded-md text-sm font-medium transition-colors',
              selectedStatus === ''
                ? 'bg-black text-white'
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
            ]"
          >
            All
          </button>
          <button
            v-for="status in statusOptions"
            :key="status"
            @click="handleStatusFilterChange(status)"
            :class="[
              'px-3 py-1 rounded-md text-sm font-medium transition-colors',
              selectedStatus === status
                ? 'bg-black text-white'
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
            ]"
          >
            {{ status }}
          </button>
        </div>
      </div>

      <div class="flex-1 flex flex-col overflow-hidden">
        <div v-if="loading" class="flex items-center justify-center h-full">
          <div class="text-gray-500">Loading...</div>
        </div>
        <div v-else-if="applications.length === 0" class="flex items-center justify-center h-full">
          <div class="text-gray-500 text-center">
            <p class="text-lg">No applications yet</p>
            <p class="text-sm mt-1">Click "Add Application" to get started</p>
          </div>
        </div>
        <div v-else class="flex-1 flex flex-col overflow-hidden">
          <div v-if="hasSelection" class="px-6 py-3 bg-blue-50 border-b flex items-center justify-between">
            <div class="flex items-center gap-4">
              <input
                type="checkbox"
                :checked="isAllSelected"
                @change="toggleSelectAll"
                class="w-4 h-4 cursor-pointer"
              />
              <span class="text-sm font-medium text-gray-700">
                {{ selectedIds.size }} selected
              </span>
            </div>
            <div class="flex items-center gap-2">
              <div class="relative">
                <button
                  @click="showBatchStatusDropdown = !showBatchStatusDropdown"
                  class="px-3 py-2 border rounded-md text-sm font-medium hover:bg-gray-100 transition-colors"
                >
                  Change Status
                </button>
                <div
                  v-if="showBatchStatusDropdown"
                  class="absolute right-0 mt-2 w-40 bg-white border rounded-lg shadow-lg z-40"
                >
                  <button
                    v-for="status in statusOptions"
                    :key="status"
                    @click="handleBatchStatusChange(status)"
                    class="w-full text-left px-4 py-2 hover:bg-blue-50 transition-colors text-sm first:rounded-t-lg last:rounded-b-lg"
                  >
                    {{ status }}
                  </button>
                </div>
              </div>
              <button
                @click="handleBatchDelete"
                class="px-3 py-2 border rounded-md text-sm font-medium text-red-600 hover:bg-red-50 transition-colors"
              >
                Delete
              </button>
              <button
                @click="clearSelection"
                class="px-3 py-2 border rounded-md text-sm font-medium hover:bg-gray-100 transition-colors"
              >
                Clear
              </button>
            </div>
          </div>

          <div class="flex-1 overflow-x-auto" @click.self="closeActionDropdown()">
            <table class="w-full border-collapse">
              <thead class="sticky top-0 bg-gray-50 border-b">
                <tr>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-8">
                    <input
                      type="checkbox"
                      :checked="isAllSelected"
                      @change="toggleSelectAll"
                      class="w-4 h-4 cursor-pointer"
                    />
                  </th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-32">Company</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-40">Position</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-40">Location</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-40">Salary Range</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-32">Status</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-32">Applied Date</th>
                  <th class="px-6 py-3 text-left text-sm font-semibold text-gray-700 w-20">Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="app in applications" :key="app.id" class="border-b hover:bg-gray-50 transition-colors">
                  <td class="px-6 py-4 text-sm">
                    <input
                      type="checkbox"
                      :checked="selectedIds.has(app.id)"
                      @change="toggleSelectId(app.id)"
                      class="w-4 h-4 cursor-pointer"
                    />
                  </td>
                  <td class="px-6 py-4 text-sm font-medium text-gray-900">{{ app.companyName }}</td>
                  <td class="px-6 py-4 text-sm text-gray-600">{{ app.positionTitle }}</td>
                  <td class="px-6 py-4 text-sm text-gray-600">{{ app.location || '-' }}</td>
                  <td class="px-6 py-4 text-sm text-gray-600">{{ app.salaryRange || '-' }}</td>
                  <td class="px-6 py-4 text-sm">
                    <select
                      :value="app.status || 'Wishlist'"
                      @change="(e) => handleStatusChange(app.id, (e.target as HTMLSelectElement).value)"
                      class="px-2 py-1 border rounded-md text-sm outline-none hover:border-gray-400 focus:border-black"
                    >
                      <option v-for="status in statusOptions" :key="status" :value="status">
                        {{ status }}
                      </option>
                    </select>
                  </td>
                  <td class="px-6 py-4 text-sm text-gray-600">
                    {{ app.appliedDate ? new Date(app.appliedDate).toLocaleDateString() : '-' }}
                  </td>
                  <td class="px-6 py-4 text-sm relative">
                    <button
                      @click="toggleActionDropdown(app.id)"
                      class="p-2 hover:bg-gray-200 rounded transition-colors"
                      title="Actions"
                    >
                      <MoreVertical24Filled class="w-4 h-4" />
                    </button>
                    <div
                      v-if="openActionDropdown === app.id"
                      class="absolute right-0 mt-0 w-48 bg-white border rounded-lg shadow-lg z-40"
                    >
                      <button
                        @click="openEditForm(app); closeActionDropdown()"
                        class="w-full text-left px-4 py-2 hover:bg-gray-100 transition-colors flex items-center gap-2 first:rounded-t-lg"
                      >
                        <Edit24Filled class="w-4 h-4" />
                        <span>Edit</span>
                      </button>
                      <button
                        @click="handleDelete(app.id); closeActionDropdown()"
                        class="w-full text-left px-4 py-2 hover:bg-red-50 transition-colors text-red-600 flex items-center gap-2 last:rounded-b-lg"
                      >
                        <Delete24Filled class="w-4 h-4" />
                        <span>Delete</span>
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="px-6 py-4 border-t bg-white flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div class="text-sm text-gray-600">
                Showing {{ offset + 1 }} to {{ Math.min(offset + pageSize, totalCount) }} of {{ totalCount }} applications
              </div>
              <div class="flex items-center gap-2">
                <label class="text-sm text-gray-600">Per page:</label>
                <select
                  :value="pageSize"
                  @change="(e) => changePageSize(Number((e.target as HTMLSelectElement).value))"
                  class="px-2 py-1 border rounded-md text-sm outline-none"
                >
                  <option value="5">5</option>
                  <option value="10">10</option>
                  <option value="25">25</option>
                  <option value="50">50</option>
                </select>
              </div>
            </div>

            <div class="flex items-center gap-2">
              <button
                @click="changePage(currentPage - 1)"
                :disabled="currentPage === 1 || loading"
                class="px-3 py-1 border rounded-md text-sm font-medium disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
              >
                Previous
              </button>
              <div class="flex items-center gap-1">
                <button
                  v-for="page in Math.min(5, totalPages)"
                  :key="page"
                  @click="changePage(page)"
                  :class="[
                    'px-3 py-1 rounded-md text-sm font-medium transition-colors',
                    page === currentPage
                      ? 'bg-black text-white'
                      : 'border hover:bg-gray-50'
                  ]"
                >
                  {{ page }}
                </button>
                <span v-if="totalPages > 5" class="px-2 text-sm text-gray-600">...</span>
              </div>
              <button
                @click="changePage(currentPage + 1)"
                :disabled="currentPage === totalPages || loading"
                class="px-3 py-1 border rounded-md text-sm font-medium disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
              >
                Next
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <Transition name="modal">
      <div v-if="showModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
          <div class="p-6">
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-xl font-bold">{{ editingId ? 'Edit Application' : 'New Application' }}</h2>
              <button
                @click="showModal = false"
                class="text-gray-500 hover:text-gray-700 text-2xl"
              >
                ×
              </button>
            </div>

            <Form :form-value="formValue" :on-register-validation="registerValidation">
              <div class="space-y-4">
                <Input
                  field="companyName"
                  label="Company Name"
                  placeholder="e.g., Google"
                  required
                  :min="2"
                  :max="255"
                />
                <Input
                  field="positionTitle"
                  label="Position Title"
                  placeholder="e.g., Senior Engineer"
                  required
                  :min="2"
                  :max="255"
                />
                <Input
                  field="jobUrl"
                  label="Job URL"
                  type="url"
                  placeholder="https://..."
                />
                <Input
                  field="location"
                  label="Location"
                  placeholder="e.g., San Francisco, CA"
                />
                <Input
                  field="salaryRange"
                  label="Salary Range"
                  placeholder="e.g., $150k - $200k"
                />
                <div>
                  <label class="block text-sm font-medium mb-1">Status</label>
                  <select
                    v-model="formValue.status"
                    class="w-full border rounded-md px-3 py-2 outline-none"
                  >
                    <option v-for="status in statusOptions" :key="status" :value="status">
                      {{ status }}
                    </option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium mb-1">Applied Date</label>
                  <input
                    v-model="formValue.appliedDate"
                    type="date"
                    class="w-full border rounded-md px-3 py-2 outline-none"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium mb-1">Notes</label>
                  <textarea
                    v-model="formValue.notes"
                    placeholder="Add any additional notes..."
                    class="w-full border rounded-md px-3 py-2 outline-none resize-none"
                    rows="3"
                  />
                </div>

                <div class="flex gap-2 pt-4">
                  <Button
                    :disabled="!canSubmit || loading"
                    label="Save"
                    @click="handleSubmit"
                  />
                  <button
                    @click="showModal = false"
                    class="flex-1 px-4 py-2 border rounded-md font-semibold hover:bg-gray-50 transition-colors"
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </Form>
          </div>
        </div>
      </div>
    </Transition>

    <Transition name="modal">
      <div v-if="showDeleteModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-md w-full mx-4">
          <div class="p-6">
            <h2 class="text-lg font-bold mb-2">Delete Application</h2>
            <p class="text-gray-600 mb-6">Are you sure you want to delete this application? This action cannot be undone.</p>
            <div class="flex gap-2 justify-end">
              <button
                @click="cancelDelete"
                class="px-4 py-2 border rounded-md font-semibold hover:bg-gray-50 transition-colors"
              >
                Cancel
              </button>
              <button
                @click="confirmDelete"
                :disabled="loading"
                class="px-4 py-2 bg-red-600 text-white rounded-md font-semibold hover:bg-red-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>

    <Transition name="modal">
      <div v-if="showBatchDeleteModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-md w-full mx-4">
          <div class="p-6">
            <h2 class="text-lg font-bold mb-2">Delete Applications</h2>
            <p class="text-gray-600 mb-6">Are you sure you want to delete {{ selectedIds.size }} application(s)? This action cannot be undone.</p>
            <div class="flex gap-2 justify-end">
              <button
                @click="cancelBatchDelete"
                class="px-4 py-2 border rounded-md font-semibold hover:bg-gray-50 transition-colors"
              >
                Cancel
              </button>
              <button
                @click="confirmBatchDelete"
                :disabled="loading"
                class="px-4 py-2 bg-red-600 text-white rounded-md font-semibold hover:bg-red-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Delete All
              </button>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
