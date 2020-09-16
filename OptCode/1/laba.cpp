#include <iostream>
#include <sstream>
#include <string>

#include <gcc-plugin.h>
#include <plugin-version.h>
#include <coretypes.h>
#include <tree-pass.h>
#include <context.h>
#include <basic-block.h>
#include <tree.h>
#include <tree-ssa-alias.h>
#include <gimple-expr.h>
#include <gimple.h>
#include <gimple-ssa.h>
#include <tree-phinodes.h>
#include <tree-ssa-operands.h>
#include <ssa-iterators.h>
#include <gimple-iterator.h>


int plugin_is_GPL_compatible;

void handle_tree(tree t) {
    switch (TREE_CODE(TREE_TYPE(t))) {
        case INTEGER_TYPE:
            std::cout << "integer ";
            break;
        case VOID_TYPE:
            std::cout << "void ";
            break;
        case REAL_TYPE:
            std::cout << "real ";
            break;
        case POINTER_TYPE:
            std::cout << "pointer ";
            break;
    }
    switch (TREE_CODE(t)) {
        case INTEGER_CST:
            std::cout << TREE_INT_CST_LOW(t);
            break;
        case REAL_CST:
            std::cout << "REAL_CST";
            break;
        case COMPLEX_CST:
            std::cout << "COMPLEX_CST";
            break;
        case STRING_CST:
            std::cout << TREE_STRING_POINTER(t);
            break;
        case LABEL_DECL:
            std::cout << (DECL_NAME(t) ? IDENTIFIER_POINTER(DECL_NAME(t)) : "unknown_label") << ":";
            break;
        case FIELD_DECL:
            std::cout << (DECL_NAME(t) ? IDENTIFIER_POINTER(DECL_NAME(t)) : "unknown_field");
            break;
        case VAR_DECL:
            std::cout << (DECL_NAME(t) ? IDENTIFIER_POINTER(DECL_NAME(t)) : "unknown_var");
            break;
        case CONST_DECL:
            std::cout << (DECL_NAME(t) ? IDENTIFIER_POINTER(DECL_NAME(t)) : "unknown_const");
            break;
        case COMPONENT_REF:
            handle_tree(TREE_OPERAND(t, 0));
            std::cout << "->";
            handle_tree(TREE_OPERAND(t, 1));
            break;
        case ARRAY_REF:
            handle_tree(TREE_OPERAND(t, 0));
            std::cout << "[";
            handle_tree(TREE_OPERAND(t, 1));
            std::cout << "]";
            break;
        case INDIRECT_REF:
            std::cout << "*";
            handle_tree(TREE_OPERAND(t, 0));
            break;
        case CONSTRUCTOR:
            std::cout << "constructor";
            break;
        case ADDR_EXPR:
            std::cout << "&";
            handle_tree(TREE_OPERAND(t, 0));
            break;
        case MEM_REF:
            std::cout << "MEM_REF(";
            handle_tree(TREE_OPERAND(t, 1));
            std::cout << ", ";
            handle_tree(TREE_OPERAND(t, 0));
            std::cout << ")";
            break;
        case SSA_NAME: {
            gimple stmt = SSA_NAME_DEF_STMT(t);
            if (gimple_code(stmt) == GIMPLE_PHI) {
                std::cout << "("
                          << (SSA_NAME_IDENTIFIER(t) ? IDENTIFIER_POINTER(SSA_NAME_IDENTIFIER(t)) : "unknown_name")
                          << "_" << SSA_NAME_VERSION(t);
                std::cout << " = GIMPLE_PHI(";
                for (unsigned int i = 0; i < gimple_phi_num_args(stmt); i++) {
                    handle_tree(gimple_phi_arg(stmt, i)->def);
                    if (i != gimple_phi_num_args(stmt) - 1) {
                        std::cout << ", ";
                    }
                }
                std::cout << "))";
            } else {
                std::cout << (SSA_NAME_IDENTIFIER(t) ? IDENTIFIER_POINTER(SSA_NAME_IDENTIFIER(t)) : "unknown_name")
                          << "_" << SSA_NAME_VERSION(t);
            }
            break;
        }
    }
}

void handle_operator(enum tree_code code) {
    switch (code) {
        case POINTER_PLUS_EXPR:
        case PLUS_EXPR:
            std::cout << "+";
            break;
        case NEGATE_EXPR:
        case MINUS_EXPR:
            std::cout << "-";
            break;
        case MULT_EXPR:
            std::cout << "*";
            break;
        case TRUNC_DIV_EXPR:
        case CEIL_DIV_EXPR:
        case FLOOR_DIV_EXPR:
        case ROUND_DIV_EXPR:
        case EXACT_DIV_EXPR:
        case RDIV_EXPR:
            std::cout << "/";
            break;
        case LSHIFT_EXPR:
            std::cout << "<<";
            break;
        case RSHIFT_EXPR:
            std::cout << ">>";
            break;
        case BIT_IOR_EXPR:
            std::cout << "|";
            break;
        case BIT_XOR_EXPR:
            std::cout << "^";
            break;
        case BIT_AND_EXPR:
            std::cout << "&";
            break;
        case BIT_NOT_EXPR:
            std::cout << "!";
            break;
        case TRUTH_ANDIF_EXPR:
        case TRUTH_AND_EXPR:
            std::cout << "&&";
            break;
        case TRUTH_ORIF_EXPR:
        case TRUTH_OR_EXPR:
            std::cout << "||";
            break;
        case TRUTH_XOR_EXPR:
            std::cout << "^^";
            break;
        case TRUTH_NOT_EXPR:
            std::cout << "!";
            break;
        case LT_EXPR:
        case UNLT_EXPR:
            std::cout << "<";
            break;
        case LE_EXPR:
        case UNLE_EXPR:
            std::cout << "<=";
            break;
        case GT_EXPR:
        case UNGT_EXPR:
            std::cout << ">";
            break;
        case GE_EXPR:
        case UNGE_EXPR:
            std::cout << ">=";
            break;
        case EQ_EXPR:
        case UNEQ_EXPR:
            std::cout << "==";
            break;
        case NE_EXPR:
        case LTGT_EXPR:
            std::cout << "!=";
            break;
        default:
            break;
    }
}

void handle_statements(basic_block bb) {

    for (gimple_stmt_iterator gsi = gsi_start_bb(bb); !gsi_end_p(gsi); gsi_next(&gsi)) {

        gimple stmt = gsi_stmt(gsi);

        switch (gimple_code(stmt)) {
            case GIMPLE_ASSIGN: {
                std::cout << "        statement: GIMPLE_ASSIGN (" << GIMPLE_ASSIGN << ") { ";
                switch (gimple_num_ops(stmt)) {
                    case 2:
                        handle_tree(gimple_assign_lhs(stmt));
                        std::cout << " = ";
                        handle_tree(gimple_assign_rhs1(stmt));
                        break;
                    case 3:
                        handle_tree(gimple_assign_lhs(stmt));
                        std::cout << " = ";
                        handle_tree(gimple_assign_rhs1(stmt));
                        std::cout << " ";
                        handle_operator(gimple_assign_rhs_code(stmt));
                        std::cout << " ";
                        handle_tree(gimple_assign_rhs2(stmt));
                        break;
                }
                std::cout << " }" << std::endl;
                break;
            }
            case GIMPLE_CALL: {
                std::cout << "        statement: GIMPLE_CALL (" << GIMPLE_CALL << ") { ";
                tree lhs = gimple_call_lhs(stmt);
                if (lhs) {
                    handle_tree(lhs);
                    printf(" = ");
                }
                std::cout << fndecl_name(gimple_call_fndecl(stmt)) << "(";
                for (unsigned int i = 0; i < gimple_call_num_args(stmt); i++) {
                    handle_tree(gimple_call_arg(stmt, i));
                    if (i != gimple_call_num_args(stmt) - 1) {
                        std::cout << ", ";
                    }
                }
                std::cout << ")";
                std::cout << " }" << std::endl;
                break;
            }
            case GIMPLE_COND: {
                std::cout << "        statement: GIMPLE_COND (" << GIMPLE_COND << ") { ";
                handle_tree(gimple_cond_lhs(stmt));
                std::cout << " ";
                handle_operator(gimple_assign_rhs_code(stmt));
                std::cout << " ";
                handle_tree(gimple_cond_rhs(stmt));
                std::cout << " }" << std::endl;
                break;
            }
            case GIMPLE_LABEL: {
                std::cout << "        statement: GIMPLE_LABEL (" << GIMPLE_LABEL << ")" << std::endl;
                break;
            }
            case GIMPLE_RETURN: {
                std::cout << "        statement: GIMPLE_RETURN (" << GIMPLE_RETURN << ")" << std::endl;
                break;
            }
        }
    }

}

void handle_basic_block(basic_block bb) {

    std::cout << "    basic block: ";
    edge e;
    edge_iterator it;
    std::cout << "(";
    int first = 1;
    FOR_EACH_EDGE(e, it, bb->preds)
    {
        first = 0;
        std::cout << e->src->index;
    }
    std::cout << ") -> (" << bb->index << ") -> (";
    first = 1;
    FOR_EACH_EDGE(e, it, bb->succs)
    {
        first = 0;
        std::cout << e->dest->index;
    }
    std::cout << ")";

}

static struct plugin_info lab1_plugin_info = {"1.0", ""};

static const struct pass_data my_pass_data = {
        GIMPLE_PASS,
        "lab1",
        OPTGROUP_NONE,
        TV_NONE,
        PROP_gimple_any,
        0,
        0,
        0,
        0,
};

struct pass : gimple_opt_pass {
    pass(gcc::context *ctx) : gimple_opt_pass(my_pass_data, ctx) {}

    virtual unsigned int execute(function *f) override {
        std::cout << "function: " << function_name(f) << " {" << std::endl;
        basic_block bb;
        FOR_EACH_BB_FN(bb, f)
        {
            handle_basic_block(bb);
            std::cout << " {" << std::endl;
            handle_statements(bb);
            std::cout << "    }" << std::endl;
        }
        std::cout << "}" << std::endl << std::endl;
        return 0;
    };
};

static struct register_pass_info pass_info = {
        new pass(g),
        "ssa",
        1,
        PASS_POS_INSERT_AFTER
};

int plugin_init(struct plugin_name_args *args, struct plugin_gcc_version *version) {

    if (!plugin_default_version_check(version, &gcc_version)) return 1;

    register_callback(args->base_name, PLUGIN_INFO, NULL, &lab1_plugin_info);
    register_callback(args->base_name, PLUGIN_PASS_MANAGER_SETUP, NULL, &pass_info);

    return 0;

}
