vim.g.mapleader = ' '

vim.keymap.set('n','<leader>f', '<cmd>Telescope find_files<cr>', { desc = "Search Files"})
vim.keymap.set('n','<leader>F', '<cmd>NvimTreeToggle<cr>', { desc = "Toggle File Tree"})

-- Window controls
vim.keymap.set('n','ws', '<cmd>split<cr>', { desc = 'HSplit'})
vim.keymap.set('n','wv', '<cmd>vsplit<cr>', { desc = 'VSplit'} )
vim.keymap.set('n','wj', '<cmd>wincmd j<cr>', { desc = 'Move down' })
vim.keymap.set('n','wk', '<cmd>wincmd k<cr>', { desc = 'Move up' })
vim.keymap.set('n','wh', '<cmd>wincmd h<cr>', { desc = 'Move left' })
vim.keymap.set('n','wl', '<cmd>wincmd l<cr>', { desc = 'Move right' })

vim.keymap.set("n", "w=", function()
     return "<cmd>res +" .. (vim.v.count == 0 and "1" or vim.v.count) .. "<cr>" end, {expr = true}, { desc = "Widen"})

vim.keymap.set("n", "w-", function()
     return "<cmd>res -" .. (vim.v.count == 0 and "1" or vim.v.count) .. "<cr>" end, {expr = true}, { desc = "Squish"})

vim.keymap.set("n", "w[", function()
     return "<cmd>vert res -" .. (vim.v.count == 0 and "1" or vim.v.count) .. "<cr>" end, {expr = true}, { desc = "Stretch"})

vim.keymap.set("n", "w]", function()
     return "<cmd>vert res +" .. (vim.v.count == 0 and "1" or vim.v.count) .. "<cr>" end, {expr = true}, { desc = "Squash"})

-- Tab controls
vim.keymap.set('n','<leader>tt', "<cmd>tabnew<cr>", { desc = "New tab" })
vim.keymap.set('n', '<leader>tq', "<cmd>tabclose<cr>", { desc = "Close tab" })
vim.keymap.set('n', ',',"<cmd>tabprevious<cr>", { desc = "Tab left"})
vim.keymap.set('n', '.',"<cmd>tabnext<cr>", { desc = "Tab right"})
vim.keymap.set('n', '<leader>t.', '<cmd>+tabmove<cr>', { desc = "Move right"})
vim.keymap.set('n', '<leader>t,', '<cmd>-tabmove<cr>', { desc = "Move left"})

-- Buffer controls
vim.keymap.set('n', '>', '<cmd>+buffer<cr>', { desc = "Buffer right"})
vim.keymap.set('n', '<', '<cmd>-buffer<cr>', { desc = "Buffer left"})

-- Tab numbers
vim.keymap.set('n', '<leader>t1', "<cmd>1tabnext<cr>", { desc = "Tab 1" })
vim.keymap.set('n', '<leader>t2', "<cmd>2tabnext<cr>", { desc = "Tab 2" })
vim.keymap.set('n', '<leader>t3', "<cmd>3tabnext<cr>", { desc = "Tab 3" })
vim.keymap.set('n', '<leader>t4', "<cmd>4tabnext<cr>", { desc = "Tab 4" })
vim.keymap.set('n', '<leader>t5', "<cmd>5tabnext<cr>", { desc = "Tab 5" })
vim.keymap.set('n', '<leader>t6', "<cmd>6tabnext<cr>", { desc = "Tab 6" })
vim.keymap.set('n', '<leader>t7', "<cmd>7tabnext<cr>", { desc = "Tab 7" })
vim.keymap.set('n', '<leader>t8', "<cmd>8tabnext<cr>", { desc = "Tab 8" })
vim.keymap.set('n', '<leader>t9', "<cmd>9tabnext<cr>", { desc = "Tab 9" })
vim.keymap.set('n', '<leader>t0', "<cmd>10tabnext<cr>", { desc = "Tab 10" })

-- CMake
vim.keymap.set('n', '<leader>cg', "<cmd>CMakeGenerate<cr>", { desc = "Generate build files"})
vim.keymap.set('n', '<leader>cb', "<cmd>CMakeBuild<cr>", { desc = "Build project"})
vim.keymap.set('n', '<leader>cr', "<cmd>CMakeRun<cr>", { desc = "Run project"})
vim.keymap.set('n', '<leader>cs', "<cmd>CMakeStop<cr>", { desc = "Stop running project"})
vim.keymap.set('n', '<leader>cf', "<cmd>Telescope cmake_tools sources<cr>", { desc = "Show project files"})
