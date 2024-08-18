// Generated from /Users/klestiselimaj/Developer/RustRoverProjects/wyst/Wyst.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class WystLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, WS=9, 
		NUMBER=10, HEX=11, IDENTIFIER=12, MATH=13, STRING=14;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "WS", 
			"NUMBER", "HEX", "IDENTIFIER", "MATH", "ESC", "STRING"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "','", "')'", "'='", "'%'", "';'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "WS", "NUMBER", 
			"HEX", "IDENTIFIER", "MATH", "STRING"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public WystLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Wyst.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u000ea\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0004\b1\b\b\u000b"+
		"\b\f\b2\u0001\b\u0001\b\u0001\t\u0001\t\u0005\t9\b\t\n\t\f\t<\t\t\u0001"+
		"\t\u0003\t?\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005\nE\b\n\n\n\f\nH"+
		"\t\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000bN\b\u000b"+
		"\n\u000b\f\u000bQ\t\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0005\u000e[\b\u000e\n\u000e\f\u000e^\t"+
		"\u000e\u0001\u000e\u0001\u000e\u0000\u0000\u000f\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u0000\u001d\u000e\u0001\u0000\t\u0003"+
		"\u0000\t\n\r\r  \u0001\u000019\u0001\u000009\u0003\u000009AFaf\u0003\u0000"+
		"AZ__az\u0003\u000009AZaz\u0003\u0000*+--//\u0003\u0000\"\"\'\'\\\\\u0002"+
		"\u0000\"\"\\\\g\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001"+
		"\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001"+
		"\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000"+
		"\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000"+
		"\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000"+
		"\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000"+
		"\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0001\u001f\u0001\u0000\u0000\u0000\u0003!\u0001\u0000\u0000\u0000"+
		"\u0005#\u0001\u0000\u0000\u0000\u0007%\u0001\u0000\u0000\u0000\t\'\u0001"+
		"\u0000\u0000\u0000\u000b)\u0001\u0000\u0000\u0000\r+\u0001\u0000\u0000"+
		"\u0000\u000f-\u0001\u0000\u0000\u0000\u00110\u0001\u0000\u0000\u0000\u0013"+
		">\u0001\u0000\u0000\u0000\u0015@\u0001\u0000\u0000\u0000\u0017I\u0001"+
		"\u0000\u0000\u0000\u0019R\u0001\u0000\u0000\u0000\u001bT\u0001\u0000\u0000"+
		"\u0000\u001dW\u0001\u0000\u0000\u0000\u001f \u0005(\u0000\u0000 \u0002"+
		"\u0001\u0000\u0000\u0000!\"\u0005,\u0000\u0000\"\u0004\u0001\u0000\u0000"+
		"\u0000#$\u0005)\u0000\u0000$\u0006\u0001\u0000\u0000\u0000%&\u0005=\u0000"+
		"\u0000&\b\u0001\u0000\u0000\u0000\'(\u0005%\u0000\u0000(\n\u0001\u0000"+
		"\u0000\u0000)*\u0005;\u0000\u0000*\f\u0001\u0000\u0000\u0000+,\u0005{"+
		"\u0000\u0000,\u000e\u0001\u0000\u0000\u0000-.\u0005}\u0000\u0000.\u0010"+
		"\u0001\u0000\u0000\u0000/1\u0007\u0000\u0000\u00000/\u0001\u0000\u0000"+
		"\u000012\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u000023\u0001\u0000"+
		"\u0000\u000034\u0001\u0000\u0000\u000045\u0006\b\u0000\u00005\u0012\u0001"+
		"\u0000\u0000\u00006:\u0007\u0001\u0000\u000079\u0007\u0002\u0000\u0000"+
		"87\u0001\u0000\u0000\u00009<\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000"+
		"\u0000:;\u0001\u0000\u0000\u0000;?\u0001\u0000\u0000\u0000<:\u0001\u0000"+
		"\u0000\u0000=?\u00050\u0000\u0000>6\u0001\u0000\u0000\u0000>=\u0001\u0000"+
		"\u0000\u0000?\u0014\u0001\u0000\u0000\u0000@A\u00050\u0000\u0000AB\u0005"+
		"x\u0000\u0000BF\u0001\u0000\u0000\u0000CE\u0007\u0003\u0000\u0000DC\u0001"+
		"\u0000\u0000\u0000EH\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000"+
		"FG\u0001\u0000\u0000\u0000G\u0016\u0001\u0000\u0000\u0000HF\u0001\u0000"+
		"\u0000\u0000IO\u0007\u0004\u0000\u0000JN\u0007\u0005\u0000\u0000KL\u0005"+
		":\u0000\u0000LN\u0005:\u0000\u0000MJ\u0001\u0000\u0000\u0000MK\u0001\u0000"+
		"\u0000\u0000NQ\u0001\u0000\u0000\u0000OM\u0001\u0000\u0000\u0000OP\u0001"+
		"\u0000\u0000\u0000P\u0018\u0001\u0000\u0000\u0000QO\u0001\u0000\u0000"+
		"\u0000RS\u0007\u0006\u0000\u0000S\u001a\u0001\u0000\u0000\u0000TU\u0005"+
		"\\\u0000\u0000UV\u0007\u0007\u0000\u0000V\u001c\u0001\u0000\u0000\u0000"+
		"W\\\u0005\"\u0000\u0000X[\u0003\u001b\r\u0000Y[\b\b\u0000\u0000ZX\u0001"+
		"\u0000\u0000\u0000ZY\u0001\u0000\u0000\u0000[^\u0001\u0000\u0000\u0000"+
		"\\Z\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000]_\u0001\u0000\u0000"+
		"\u0000^\\\u0001\u0000\u0000\u0000_`\u0005\"\u0000\u0000`\u001e\u0001\u0000"+
		"\u0000\u0000\t\u00002:>FMOZ\\\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}