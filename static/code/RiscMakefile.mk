all:
        riscv64-unknown-elf-gcc -c main.s
        riscv64-unknown-elf-ld main.o
qemugdb: all
        qemu-riscv64 -g 9000 a.out
gdb:
        riscv64-unknown-elf-gdb

