-- this file is for configuring the lsp for the monkey lang


local on_attach = function(client, bufnr)
    local buf_map = function(bufnr, mode, lhs, rhs, opts)
        opts = vim.tbl_extend('force', {noremap = true, silent = true}, opts or {})
        vim.api.nvim_buf_set_keymap(bufnr, mode, lhs, rhs, opts)
    end

    -- Mappings for LSP functionality
    buf_map(bufnr, 'n', 'gd', '<cmd>lua vim.lsp.buf.definition()<CR>')
    buf_map(bufnr, 'n', 'K', '<cmd>lua vim.lsp.buf.hover()<CR>')
    buf_map(bufnr, 'n', 'gi', '<cmd>lua vim.lsp.buf.implementation()<CR>')
    buf_map(bufnr, 'n', '<C-k>', '<cmd>lua vim.lsp.buf.signature_help()<CR>')
    buf_map(bufnr, 'n', '<leader>rn', '<cmd>lua vim.lsp.buf.rename()<CR>')
    buf_map(bufnr, 'n', 'gr', '<cmd>lua vim.lsp.buf.references()<CR>')
    buf_map(bufnr, 'n', '<leader>ca', '<cmd>lua vim.lsp.buf.code_action()<CR>')
end

local cli = vim.lsp.start_client{
    name = "shitty lsp by him",
    cmd = {"/home/boburmirzoalivobjonov/Programming/Projects/LSP/bin/main"},
    on_attach = on_attach
}

if cli then
    vim.notify "hey shitty lsp started"
end

vim.api.nvim_create_autocmd("FileType", {
    pattern = "monkey",
    callback = function ()
       vim.lsp.buf_attach_client(0, cli)
    end
})
