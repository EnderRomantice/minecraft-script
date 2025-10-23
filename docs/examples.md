# Minecraft Script 使用示例与流程说明

## 基本使用流程

1. **安装**
   ```bash
   go install github.com/minecraft-script/cmd/minecraftscript@latest
   ```

2. **创建脚本文件**
   创建一个扩展名为 `.mcs` 的文件，例如 `build_house.mcs`

3. **运行脚本**
   ```bash
   minecraftscript run build_house.mcs
   ```

4. **查看生成的命令**
   命令将输出到控制台或指定的输出文件

## 语法示例

### 1. 变量定义与使用

```
# 定义坐标变量
start_pos = (0, 64, 0)
end_pos = (10, 74, 10)

# 使用变量创建一个石头立方体
fill(start_pos, end_pos, "stone")
```

生成的 Minecraft 命令：
```
/fill 0 64 0 10 74 10 stone
```

### 2. 创建简单房屋

```
# 定义房屋坐标
floor_start = (0, 64, 0)
floor_end = (10, 64, 10)
wall_height = 5
roof_height = 7

# 创建地板
fill(floor_start, floor_end, "oak_planks")

# 创建墙壁 - 前后墙
fill((floor_start[0], floor_start[1] + 1, floor_start[2]), (floor_end[0], floor_start[1] + wall_height, floor_start[2]), "oak_log")
fill((floor_start[0], floor_start[1] + 1, floor_end[2]), (floor_end[0], floor_start[1] + wall_height, floor_end[2]), "oak_log")

# 创建墙壁 - 左右墙
fill((floor_start[0], floor_start[1] + 1, floor_start[2]), (floor_start[0], floor_start[1] + wall_height, floor_end[2]), "oak_log")
fill((floor_end[0], floor_start[1] + 1, floor_start[2]), (floor_end[0], floor_start[1] + wall_height, floor_end[2]), "oak_log")

# 创建屋顶
fill((floor_start[0], floor_start[1] + roof_height, floor_start[2]), (floor_end[0], floor_start[1] + roof_height, floor_end[2]), "dark_oak_planks")

# 创建门
setblock((floor_start[0] + 5, floor_start[1] + 1, floor_start[2]), "oak_door")
```

### 3. 创建简单的农场

```
# 定义农场坐标
farm_start = (0, 64, 0)
farm_size = 10

# 创建耕地
for x in range(farm_start[0], farm_start[0] + farm_size):
    for z in range(farm_start[2], farm_start[2] + farm_size):
        setblock((x, farm_start[1], z), "farmland")
        setblock((x, farm_start[1] + 1, z), "wheat")

# 创建围栏
for x in range(farm_start[0] - 1, farm_start[0] + farm_size + 1):
    setblock((x, farm_start[1], farm_start[2] - 1), "oak_fence")
    setblock((x, farm_start[1], farm_start[2] + farm_size), "oak_fence")

for z in range(farm_start[2], farm_start[2] + farm_size):
    setblock((farm_start[0] - 1, farm_start[1], z), "oak_fence")
    setblock((farm_start[0] + farm_size, farm_start[1], z), "oak_fence")
```

## 编译与执行流程

Minecraft Script 的编译与执行流程如下：

1. **词法分析**
   - 输入：源代码文本
   - 处理：将源代码转换为词法单元序列
   - 输出：词法单元序列

2. **语法分析**
   - 输入：词法单元序列
   - 处理：构建抽象语法树
   - 输出：抽象语法树

3. **代码生成**
   - 输入：抽象语法树
   - 处理：遍历语法树，生成 Minecraft 命令
   - 输出：Minecraft 命令序列

4. **执行**
   - 输入：Minecraft 命令序列
   - 处理：在 Minecraft 游戏中执行命令
   - 输出：游戏中的构建结果

## 调试技巧

1. **查看词法分析结果**
   ```bash
   minecraftscript debug --tokens script.mcs
   ```

2. **查看语法树**
   ```bash
   minecraftscript debug --ast script.mcs
   ```

3. **仅生成命令而不执行**
   ```bash
   minecraftscript compile script.mcs > commands.txt
   ```

4. **逐步执行命令**
   ```bash
   minecraftscript run --step script.mcs
   ```

## 常见错误与解决方案

1. **语法错误**
   - 错误：`Syntax error: unexpected token at line X`
   - 解决：检查该行的语法，确保括号匹配、逗号使用正确

2. **未定义变量**
   - 错误：`Undefined variable: X`
   - 解决：确保在使用变量前已定义

3. **类型错误**
   - 错误：`Type error: expected vector, got string`
   - 解决：确保变量类型与使用场景匹配

4. **命令生成失败**
   - 错误：`Failed to generate command for expression at line X`
   - 解决：检查表达式是否符合 Minecraft 命令的要求